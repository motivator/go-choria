package client

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/google/go-cmp/cmp"
	"github.com/tidwall/gjson"

	"github.com/choria-io/go-choria/config"
	"github.com/choria-io/go-choria/providers/discovery/broadcast"
	"github.com/choria-io/go-choria/providers/discovery/puppetdb"

	"github.com/choria-io/go-choria/choria"
	cclient "github.com/choria-io/go-choria/client/client"
	"github.com/choria-io/go-choria/protocol"
	"github.com/choria-io/go-choria/providers/agent/mcorpc"
	addl "github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
	"github.com/choria-io/go-choria/srvcache"

	"github.com/sirupsen/logrus"
)

type ChoriaFramework interface {
	Logger(string) *logrus.Entry
	Configuration() *config.Config
	NewMessage(payload string, agent string, collective string, msgType string, request *choria.Message) (msg *choria.Message, err error)
	NewReplyFromTransportJSON(payload []byte, skipvalidate bool) (msg protocol.Reply, err error)
	NewTransportFromJSON(data string) (message protocol.TransportMessage, err error)
	MiddlewareServers() (servers srvcache.Servers, err error)
	NewConnector(ctx context.Context, servers func() (srvcache.Servers, error), name string, logger *logrus.Entry) (conn choria.Connector, err error)
	NewRequestID() (string, error)
	Certname() string
	PQLQueryCertNames(query string) ([]string, error)
}

// RPC is a MCollective compatible RPC client
type RPC struct {
	fw   ChoriaFramework
	opts *RequestOptions
	log  *logrus.Entry
	cfg  *config.Config
	dm   string

	agent string

	mu *sync.Mutex

	ddl *addl.DDL

	// used for testing only
	cl ChoriaClient
}

// RPCRequest is a basic RPC request
type RPCRequest struct {
	Agent  string          `json:"agent"`
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

// RPCReply is a basic RPC reply
type RPCReply struct {
	Statuscode mcorpc.StatusCode `json:"statuscode"`
	Statusmsg  string            `json:"statusmsg"`
	Data       json.RawMessage   `json:"data"`
	Sender     string            `json:"-"`
	Time       time.Time         `json:"-"`
}

// MatchExpr determines if the Reply  matches expression q using the expr format.
// The query q is expected to return a boolean type else an error will be raised
func (r *RPCReply) MatchExpr(q string, prog *vm.Program) (bool, error) {
	env := map[string]interface{}{
		"msg":            r.Statusmsg,
		"code":           int(r.Statuscode),
		"data":           r.lookup,
		"ok":             r.isOK,
		"aborted":        r.isAborted,
		"invalid_data":   r.isInvalidData,
		"missing_data":   r.isMissingData,
		"unknown_action": r.isUnknownAction,
		"unknown_error":  r.isUnknownError,
		"include":        r.include,
		"sender":         func() string { return r.Sender },
		"time":           func() time.Time { return r.Time },
	}

	var err error
	if prog == nil {
		prog, err = expr.Compile(q, expr.AsBool(), expr.AllowUndefinedVariables(), expr.Env(env))
		if err != nil {
			return false, err
		}
	}

	out, err := expr.Run(prog, env)
	if err != nil {
		return false, err
	}

	matched, ok := out.(bool)
	if !ok {
		return false, fmt.Errorf("match expressions should return boolean")
	}

	return matched, nil
}

func (r *RPCReply) isOK() bool {
	return r.Statuscode == mcorpc.OK
}

func (r *RPCReply) isAborted() bool {
	return r.Statuscode == mcorpc.Aborted
}

func (r *RPCReply) isUnknownAction() bool {
	return r.Statuscode == mcorpc.UnknownAction
}

func (r *RPCReply) isMissingData() bool {
	return r.Statuscode == mcorpc.MissingData
}

func (r *RPCReply) isInvalidData() bool {
	return r.Statuscode == mcorpc.InvalidData
}

func (r *RPCReply) isUnknownError() bool {
	return r.Statuscode == mcorpc.UnknownError
}

// https://github.com/tidwall/gjson/blob/master/SYNTAX.md
func (r *RPCReply) lookup(query string) interface{} {
	return gjson.GetBytes(r.Data, query).Value()
}

func (r *RPCReply) include(hay []interface{}, needle interface{}) bool {
	// gjson always turns numbers into float64
	i, ok := needle.(int)
	if ok {
		needle = float64(i)
	}

	for _, i := range hay {
		if cmp.Equal(i, needle) {
			return true
		}
	}

	return false
}

// RequestResult is the result of a request
type RequestResult interface {
	Stats() *Stats
}

// Handler is a function that should handle each reply synchronously
type Handler func(protocol.Reply, *RPCReply)

// ChoriaClient implements the connection to the Choria network
type ChoriaClient interface {
	Request(ctx context.Context, msg *choria.Message, handler cclient.Handler) (err error)
}

// Connector is a connection to the choria network
type Connector interface {
	QueueSubscribe(ctx context.Context, name string, subject string, group string, output chan *choria.ConnectorMessage) error
	Publish(msg *choria.Message) error
}

// Option configures the RPC client
type Option func(r *RPC)

// DDL supplies a DDL when creating the client thus avoiding a disk search
func DDL(d *addl.DDL) Option {
	return func(r *RPC) {
		r.ddl = d
	}
}

// DiscoveryMethod sets a specific discovery method
func DiscoveryMethod(dm string) Option {
	return func(r *RPC) {
		r.dm = dm
	}
}

// New creates a new RPC request
//
// A DDL is required when one is not given using the DDL() option as argument
// attempts will be made to find it on the file system should this fail an error will be returned
func New(fw ChoriaFramework, agent string, opts ...Option) (rpc *RPC, err error) {
	rpc = &RPC{
		fw:    fw,
		cfg:   fw.Configuration(),
		mu:    &sync.Mutex{},
		log:   fw.Logger("mcorpc"),
		agent: agent,
		dm:    fw.Configuration().DefaultDiscoveryMethod,
	}

	for _, opt := range opts {
		opt(rpc)
	}

	if rpc.ddl == nil {
		rpc.ddl, err = addl.Find(agent, rpc.cfg.LibDir)
		if err != nil {
			return nil, fmt.Errorf("could not load %s DDL: %s", agent, err)
		}
	}

	if rpc.ddl.Metadata.Name != agent {
		return nil, fmt.Errorf("the DDL does not describe the %s agent", agent)
	}

	return rpc, nil
}

func (r *RPC) setOptions(opts ...RequestOption) (err error) {
	r.opts, err = NewRequestOptions(r.fw, r.ddl)
	if err != nil {
		return err
	}

	for _, opt := range opts {
		opt(r.opts)
	}

	return nil
}

// Do performs a RPC request and optionally processes replies
//
// If a filter is supplied using the Filter() option and Targets() are not then discovery will be done for you
// using the broadcast method, should no nodes be discovered an error will be returned
func (r *RPC) Do(ctx context.Context, action string, payload interface{}, opts ...RequestOption) (RequestResult, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// we want to force the passing of options on every request
	err := r.setOptions(opts...)
	if err != nil {
		return nil, err
	}

	dctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if len(r.opts.Targets) == 0 {
		err := r.discover(ctx)
		if err != nil {
			return nil, fmt.Errorf("discovery failed: %s", err)
		}
	}

	discoveredCnt := len(r.opts.Targets)
	msg, cl, err := r.setupMessage(dctx, action, payload, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not configure message: %s", err)
	}

	if r.opts.DiscoveryEndCB != nil {
		err = r.opts.DiscoveryEndCB(discoveredCnt, len(r.opts.Targets))
		if err != nil {
			return nil, err
		}
	}

	r.opts.totalStats.Start()
	defer r.opts.totalStats.End()

	ctr := 0

	// the client is always batched, when batched mode is not request the size of
	// the batch matches the size of the total targets and during setupMessage()
	// an appropriate connection will be made
	err = InGroups(r.opts.Targets, r.opts.BatchSize, func(nodes []string) error {
		r.opts.stats = NewStats()
		r.opts.stats.SetDiscoveredNodes(nodes)
		msg.DiscoveredHosts = nodes

		r.opts.stats.Start()
		defer r.opts.totalStats.Merge(r.opts.stats)
		defer r.opts.stats.End()

		if ctr > 0 {
			err := InterruptableSleep(dctx, r.opts.BatchSleep)
			if err != nil {
				return err
			}
		}

		r.log.Debugf("Performing batched request %d for %d/%d nodes", ctr, len(nodes), len(r.opts.Targets))

		err = r.request(dctx, msg, cl)
		if err != nil {
			return err
		}

		ctr++

		return nil
	})

	r.opts.totalStats.SetAction(action)
	r.opts.totalStats.SetAgent(r.agent)

	return &RequestOptions{totalStats: r.opts.totalStats}, err
}

func (r *RPC) discover(ctx context.Context) error {
	if len(r.opts.Filter.Compound) > 0 {
		r.dm = "choria"
	}
	if r.opts.DiscoveryStartCB != nil {
		r.opts.DiscoveryStartCB()
	}

	r.opts.totalStats.StartDiscover()
	defer r.opts.totalStats.EndDiscover()

	if r.opts.Filter == nil {
		r.opts.Filter = protocol.NewFilter()
	}

	r.opts.Filter.AddAgentFilter(r.agent)

	var n []string
	var err error

	switch r.dm {
	case "choria":
		pdb := puppetdb.New(r.fw)
		n, err = pdb.Discover(ctx, puppetdb.Filter(r.opts.Filter), puppetdb.Timeout(r.opts.DiscoveryTimeout), puppetdb.Collective(r.opts.Collective))

	default:
		b := broadcast.New(r.fw)
		n, err = b.Discover(ctx, broadcast.Filter(r.opts.Filter), broadcast.Timeout(r.opts.DiscoveryTimeout), broadcast.Name(r.opts.ConnectionName), broadcast.Collective(r.opts.Collective))
	}
	if err != nil {
		return err
	}

	if len(n) == 0 {
		return fmt.Errorf("no targets were discovered")
	}

	r.opts.Targets = n

	return nil
}

func (r *RPC) setupMessage(ctx context.Context, action string, payload interface{}, opts ...RequestOption) (msg *choria.Message, cl ChoriaClient, err error) {
	pj, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, fmt.Errorf("could not encode payload: %s", err)
	}

	rpcreq := &RPCRequest{
		Agent:  r.agent,
		Action: action,
		Data:   pj,
	}

	rpcp, err := json.Marshal(rpcreq)
	if err != nil {
		return nil, nil, fmt.Errorf("could not encode request: %s", err)
	}

	msg, err = r.fw.NewMessage(string(rpcp), r.agent, r.cfg.MainCollective, "request", nil)
	if err != nil {
		return nil, nil, err
	}

	err = r.opts.ConfigureMessage(msg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not configure Message: %s", err)
	}

	cl = r.cl

	if r.cl == nil {
		if r.opts.BatchSize == len(r.opts.Targets) || !r.opts.ProcessReplies {
			cl, err = r.unbatchedClient()
			if err != nil {
				return nil, nil, err
			}
		} else {
			cl, err = r.batchedClient(ctx, msg.RequestID)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return msg, cl, err
}

func (r *RPC) unbatchedClient() (cl ChoriaClient, err error) {
	cl, err = cclient.New(
		r.fw,
		cclient.Receivers(r.opts.Workers),
		cclient.Timeout(r.opts.Timeout),
		cclient.OnPublishStart(r.opts.totalStats.StartPublish),
		cclient.OnPublishFinish(r.opts.totalStats.EndPublish),
		cclient.Name(r.opts.ConnectionName),
	)
	if err != nil {
		return nil, fmt.Errorf("could not setup client: %s", err)
	}

	return cl, nil
}

func (r *RPC) batchedClient(ctx context.Context, msgid string) (cl ChoriaClient, err error) {
	conn, err := r.connectBatchedConnection(ctx, fmt.Sprintf("%s_batched", msgid))
	if err != nil {
		return nil, fmt.Errorf("could not connect batched network connection: %s", err)
	}

	cl, err = cclient.New(
		r.fw,
		cclient.Receivers(r.opts.Workers),
		cclient.Timeout(r.opts.Timeout),
		cclient.OnPublishStart(r.opts.totalStats.StartPublish),
		cclient.OnPublishFinish(r.opts.totalStats.EndPublish),
		cclient.Connection(conn),
		cclient.Name(r.opts.ConnectionName),
	)
	if err != nil {
		return nil, fmt.Errorf("could not set up batched client: %s", err)
	}

	return cl, nil
}

// Reset removes the cached options, any further Do() calls need to specify full options
func (r *RPC) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.opts = nil
	r.cl = nil
}

func (r *RPC) request(ctx context.Context, msg *choria.Message, cl ChoriaClient) error {
	rctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := cl.Request(rctx, msg, r.handlerFactory(rctx, cancel))
	if err != nil {
		return err
	}

	return nil
}

func (r *RPC) handlerFactory(_ context.Context, cancel func()) cclient.Handler {
	if !r.opts.ProcessReplies {
		return nil
	}

	var prog *vm.Program
	var err error

	if r.opts.ReplyExprFilter != "" {
		prog, err = expr.Compile(r.opts.ReplyExprFilter, expr.AsBool(), expr.AllowUndefinedVariables())
		if err != nil {
			r.log.Errorf("failed to compile expression: %v", err)
			prog = nil
		}
	}

	handler := func(ctx context.Context, rawmsg *choria.ConnectorMessage) {
		reply, err := r.fw.NewReplyFromTransportJSON(rawmsg.Data, false)
		if err != nil {
			r.opts.stats.FailedRequestInc()
			r.log.Errorf("Could not process a reply: %s", err)
			return
		}

		r.opts.stats.RecordReceived(reply.SenderID())

		rpcreply, err := ParseReplyData([]byte(reply.Message()))
		if err != nil {
			r.opts.stats.FailedRequestInc()
			r.log.Errorf("Could not process reply from %s: %s", reply.SenderID(), err)
			return
		}

		rpcreply.Sender = reply.SenderID()
		rpcreply.Time = reply.Time()

		if rpcreply.Statuscode == mcorpc.OK {
			r.opts.stats.PassedRequestInc()
		} else {
			r.opts.stats.FailedRequestInc()
		}

		if r.opts.Handler != nil {
			shouldShow := true
			if r.opts.ReplyExprFilter != "" {
				shouldShow, err = rpcreply.MatchExpr(r.opts.ReplyExprFilter, prog)
				if err != nil {
					r.log.Errorf("Expr filter parsing failed in reply from %s: %s", reply.SenderID(), err)
				}
			}

			if shouldShow {
				r.opts.Handler(reply, rpcreply)
			} else {
				r.opts.Handler(reply, nil)
			}
		}

		if r.opts.stats.All() {
			cancel()
			return
		}
	}

	return handler
}

func (r *RPC) connectBatchedConnection(ctx context.Context, name string) (Connector, error) {
	connector, err := r.fw.NewConnector(ctx, r.fw.MiddlewareServers, name, r.log)
	if err != nil {
		return nil, fmt.Errorf("could not create connector: %s", err)
	}

	closer := func() {
		<-ctx.Done()

		r.log.Debugf("Closing batched connection %s", name)
		connector.Close()
	}

	go closer()

	return connector, nil
}
