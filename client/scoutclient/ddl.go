// generated code; DO NOT EDIT

package scoutclient

import (
	"encoding/base64"
	"encoding/json"

	"github.com/choria-io/go-choria/providers/agent/mcorpc/ddl/agent"
)

var rawDDL = "ewogICIkc2NoZW1hIjogImh0dHBzOi8vY2hvcmlhLmlvL3NjaGVtYXMvbWNvcnBjL2RkbC92MS9hZ2VudC5qc29uIiwKICAibWV0YWRhdGEiOiB7CiAgICAibGljZW5zZSI6ICJBcGFjaGUtMi4wIiwKICAgICJhdXRob3IiOiAiUi5JLlBpZW5hYXIgPHJpcEBkZXZjby5uZXQ+IiwKICAgICJ0aW1lb3V0IjogNSwKICAgICJuYW1lIjogInNjb3V0IiwKICAgICJ2ZXJzaW9uIjogIjAuMC4xIiwKICAgICJ1cmwiOiAiaHR0cHM6Ly9jaG9yaWEuaW8iLAogICAgImRlc2NyaXB0aW9uIjogIkNob3JpYSBTY291dCBBZ2VudCBNYW5hZ2VtZW50IEFQSSIsCiAgICAicHJvdmlkZXIiOiAiZ29sYW5nIgogIH0sCiAgImFjdGlvbnMiOiBbCiAgICB7CiAgICAgICJhY3Rpb24iOiAiY2hlY2tzIiwKICAgICAgImRpc3BsYXkiOiAib2siLAogICAgICAiZGVzY3JpcHRpb24iOiAiT2J0YWluIGEgbGlzdCBvZiBjaGVja3MgYW5kIHRoZWlyIGN1cnJlbnQgc3RhdHVzIiwKICAgICAgImlucHV0Ijoge30sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgImNoZWNrcyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJEZXRhaWxzIGFib3V0IGVhY2ggY2hlY2siLAogICAgICAgICAgInR5cGUiOiAiYXJyYXkiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiQ2hlY2tzIgogICAgICAgIH0KICAgICAgfQogICAgfSwKICAgIHsKICAgICAgImFjdGlvbiI6ICJyZXN1bWUiLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImNoZWNrcyI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiQ2hlY2tzIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJDaGVjayB0byByZXN1bWUsIGVtcHR5IG1lYW5zIGFsbCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlCiAgICAgICAgfQogICAgICB9LAogICAgICAib3V0cHV0IjogewogICAgICAgICJmYWlsZWQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTGlzdCBvZiBjaGVja3MgdGhhdCBjb3VsZCBub3QgYmUgcmVzdW1lZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJGYWlsZWQiLAogICAgICAgICAgInR5cGUiOiAiYXJyYXkiCiAgICAgICAgfSwKICAgICAgICAidHJhbnNpdGlvbmVkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxpc3Qgb2YgY2hlY2tzIHRoYXQgd2VyZSByZXN1bWVkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlRyaWdnZXJlZCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9LAogICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxpc3Qgb2YgY2hlY2tzIHRoYXQgd2FzIHNraXBwZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU2tpcHBlZCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJkaXNwbGF5IjogImZhaWxlZCIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJSZXN1bWUgYWN0aXZlIGNoZWNraW5nIG9mIG9uZSBvciBtb3JlIGNoZWNrcyIKICAgIH0sCiAgICB7CiAgICAgICJhY3Rpb24iOiAibWFpbnRlbmFuY2UiLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImNoZWNrcyI6IHsKICAgICAgICAgICJwcm9tcHQiOiAiQ2hlY2tzIiwKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJDaGVjayB0byBwYXVzZSwgZW1wdHkgbWVhbnMgYWxsIiwKICAgICAgICAgICJ0eXBlIjogImFycmF5IiwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgImZhaWxlZCI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMaXN0IG9mIGNoZWNrcyB0aGF0IGNvdWxkIG5vdCBiZSBwYXVzZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiRmFpbGVkIiwKICAgICAgICAgICJ0eXBlIjogImFycmF5IgogICAgICAgIH0sCiAgICAgICAgInRyYW5zaXRpb25lZCI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMaXN0IG9mIGNoZWNrcyB0aGF0IHdlcmUgcGF1c2VkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlRyaWdnZXJlZCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9LAogICAgICAgICJza2lwcGVkIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkxpc3Qgb2YgY2hlY2tzIHRoYXQgd2FzIHNraXBwZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU2tpcHBlZCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJkaXNwbGF5IjogImZhaWxlZCIsCiAgICAgICJkZXNjcmlwdGlvbiI6ICJQYXVzZSBjaGVja2luZyBvZiBvbmUgb3IgbW9yZSBjaGVja3MiCiAgICB9LAogICAgewogICAgICAiYWN0aW9uIjogImdvc3NfdmFsaWRhdGUiLAogICAgICAiZGVzY3JpcHRpb24iOiAiUGVyZm9ybXMgYSBHb3NzIHZhbGlkYXRpb24gdXNpbmcgYSBzcGVjaWZpYyBmaWxlIiwKICAgICAgImRpc3BsYXkiOiAiZmFpbGVkIiwKICAgICAgImFnZ3JlZ2F0ZSI6IFsKICAgICAgICB7CiAgICAgICAgICAiZnVuY3Rpb24iOiAic3VtbWFyeSIsCiAgICAgICAgICAiYXJncyI6IFsKICAgICAgICAgICAgInRlc3RzIiwKICAgICAgICAgICAgeyJmb3JtYXQiOiAgIiVzIFRlc3RzIG9uICVkIG5vZGUocykifQogICAgICAgICAgXQogICAgICAgIH0sCiAgICAgICAgewogICAgICAgICAgImZ1bmN0aW9uIjogInN1bW1hcnkiLAogICAgICAgICAgImFyZ3MiOiBbCiAgICAgICAgICAgICJmYWlsdXJlcyIsCiAgICAgICAgICAgIHsiZm9ybWF0IjogICIlcyBGYWlsZWQgdGVzdCBvbiAlZCBub2RlKHMpIn0KICAgICAgICAgIF0KICAgICAgICB9LAogICAgICAgIHsKICAgICAgICAgICJmdW5jdGlvbiI6ICJzdW1tYXJ5IiwKICAgICAgICAgICJhcmdzIjogWwogICAgICAgICAgICAic3VjY2VzcyIsCiAgICAgICAgICAgIHsiZm9ybWF0IjogICIlcyBQYXNzZWQgdGVzdHMgb24gJWQgbm9kZShzKSJ9CiAgICAgICAgICBdCiAgICAgICAgfQogICAgICBdLAogICAgICAiaW5wdXQiOiB7CiAgICAgICAgImZpbGUiOiB7CiAgICAgICAgICAicHJvbXB0IjogIkdvc3MgRmlsZSIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiUGF0aCB0byB0aGUgR29zcyB2YWxpZGF0aW9uIHNwZWNpZmljYXRpb24iLAogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJtYXhsZW5ndGgiOiAyNTYsCiAgICAgICAgICAidmFsaWRhdGlvbiI6ICIuKyIsCiAgICAgICAgICAib3B0aW9uYWwiOiBmYWxzZQogICAgICAgIH0sCiAgICAgICAgInZhcnMiOiB7CiAgICAgICAgICAicHJvbXB0IjogIlZhcnMgRmlsZSIsCiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiUGF0aCB0byBhIGZpbGUgdG8gdXNlIGFzIHRlbXBsYXRlIHZhcmlhYmxlcyIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgIm1heGxlbmd0aCI6IDI1NiwKICAgICAgICAgICJ2YWxpZGF0aW9uIjogIi4rIiwKICAgICAgICAgICJvcHRpb25hbCI6IHRydWUKICAgICAgICB9CiAgICAgIH0sCiAgICAgICJvdXRwdXQiOiB7CiAgICAgICAgInRlc3RzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBudW1iZXIgb2YgdGVzdHMgdGhhdCB3ZXJlIHJ1biIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJUZXN0cyIsCiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIgogICAgICAgIH0sCiAgICAgICAgImZhaWx1cmVzIjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIlRoZSBudW1iZXIgb2YgdGVzdHMgdGhhdCBmYWlsZWQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiRmFpbGVkIFRlc3RzIiwKICAgICAgICAgICJ0eXBlIjogImludGVnZXIiCiAgICAgICAgfSwKICAgICAgICAicnVudGltZSI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgdGltZSBpdCB0b29rIHRvIHJ1biB0aGUgdGVzdHMsIGluIHNlY29uZHMiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiUnVudGltZSIsCiAgICAgICAgICAidHlwZSI6ICJpbnRlZ2VyIgogICAgICAgIH0sCiAgICAgICAgInN1Y2Nlc3MiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiSW5kaWNhdGVzIGlmIHRoZSB0ZXN0IHBhc3NlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJTdWNjZXNzIiwKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIKICAgICAgICB9LAogICAgICAgICJzdW1tYXJ5IjogewogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkEgaHVtYW4gZnJpZW5kbHkgdGVzdCByZXN1bHQiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiU3VtbWFyeSIsCiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciCiAgICAgICAgfSwKICAgICAgICAicmVzdWx0cyI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgZnVsbCB0ZXN0IHJlc3VsdHMiLAogICAgICAgICAgImRpc3BsYXlfYXMiOiAiUmVzdWx0cyIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICB7CiAgICAgICJhY3Rpb24iOiAidHJpZ2dlciIsCiAgICAgICJpbnB1dCI6IHsKICAgICAgICAiY2hlY2tzIjogewogICAgICAgICAgInByb21wdCI6ICJDaGVja3MiLAogICAgICAgICAgImRlc2NyaXB0aW9uIjogIkNoZWNrIHRvIHRyaWdnZXIsIGVtcHR5IG1lYW5zIGFsbCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIsCiAgICAgICAgICAib3B0aW9uYWwiOiB0cnVlCiAgICAgICAgfQogICAgICB9LAogICAgICAib3V0cHV0IjogewogICAgICAgICJmYWlsZWQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTGlzdCBvZiBjaGVja3MgdGhhdCBjb3VsZCBub3QgYmUgdHJpZ2dlcmVkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIkZhaWxlZCIsCiAgICAgICAgICAidHlwZSI6ICJhcnJheSIKICAgICAgICB9LAogICAgICAgICJ0cmFuc2l0aW9uZWQiOiB7CiAgICAgICAgICAiZGVzY3JpcHRpb24iOiAiTGlzdCBvZiBjaGVja3MgdGhhdCB3ZXJlIHRyaWdnZXJlZCIsCiAgICAgICAgICAiZGlzcGxheV9hcyI6ICJUcmlnZ2VyZWQiLAogICAgICAgICAgInR5cGUiOiAiYXJyYXkiCiAgICAgICAgfSwKICAgICAgICAic2tpcHBlZCI6IHsKICAgICAgICAgICJkZXNjcmlwdGlvbiI6ICJMaXN0IG9mIGNoZWNrcyB0aGF0IHdhcyBza2lwcGVkIiwKICAgICAgICAgICJkaXNwbGF5X2FzIjogIlNraXBwZWQiLAogICAgICAgICAgInR5cGUiOiAiYXJyYXkiCiAgICAgICAgfQogICAgICB9LAogICAgICAiZGlzcGxheSI6ICJmYWlsZWQiLAogICAgICAiZGVzY3JpcHRpb24iOiAiRm9yY2UgYW4gaW1tZWRpYXRlIGNoZWNrIG9mIG9uZSBvciBtb3JlIGNoZWNrcyIKICAgIH0KICBdCn0K"

// DDLBytes is the raw JSON encoded DDL file for the agent
func DDLBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(rawDDL)
}

// DDL is a parsed and loaded DDL for the agent
func DDL() (*agent.DDL, error) {
	ddlj, err := DDLBytes()
	if err != nil {
		return nil, err
	}

	ddl := &agent.DDL{}
	err = json.Unmarshal(ddlj, ddl)
	if err != nil {
		return nil, err
	}

	return ddl, nil
}
