package provtarget

import (
	"fmt"
	"sync"

	"github.com/choria-io/go-choria/provtarget/builddefaults"
	"github.com/sirupsen/logrus"

	"github.com/choria-io/go-choria/srvcache"
)

// TargetResolver is capable of resolving the target brokers for provisioning in a comma sep list
type TargetResolver interface {
	Name() string
	Targets(*logrus.Entry) []string
}

var mu = &sync.Mutex{}
var resolver = TargetResolver(builddefaults.Provider())

// RegisterTargetResolver registers a custom target resolver, else the default will be used
func RegisterTargetResolver(r TargetResolver) error {
	mu.Lock()
	defer mu.Unlock()

	resolver = r

	return nil
}

// Targets is a list of brokers to connect to
func Targets(log *logrus.Entry) ([]srvcache.Server, error) {
	mu.Lock()
	defer mu.Unlock()

	if resolver == nil {
		return []srvcache.Server{}, fmt.Errorf("no Provisioning Target Resolver registered")
	}

	s := resolver.Targets(log)

	if len(s) == 0 {
		return []srvcache.Server{}, fmt.Errorf("provisioning target plugin %s returned no servers", Name())
	}

	servers, err := srvcache.StringHostsToServers(s, "nats")
	if err != nil {
		return []srvcache.Server{}, fmt.Errorf("could not determine provisioning servers using %s provisionig target plugin: %s", Name(), err)
	}

	if len(servers) == 0 {
		return []srvcache.Server{}, fmt.Errorf("provisioning broker urls from the %s plugin were not in the valid format, 0 server:port combinations were foundin %v", Name(), s)
	}

	return servers, nil
}

// Name is the name of the plugin used
func Name() string {
	if resolver == nil {
		return "Unknown"
	}

	return resolver.Name()
}