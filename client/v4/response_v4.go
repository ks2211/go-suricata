package v4

import "github.com/ks2211/go-suricata/client"

// MemCapListResponse is message from memcap-list
type MemCapListResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MemCapShowResponse is message from memcap-show of a specific memcap
type MemCapShowResponse struct {
	Value string `json:"value"`
}

// MemCapSetResponse is message from memcap-set for a specific memcat
type MemCapSetResponse client.InformationResponse

// String is a helper method to turn go type into struct
func (m MemCapSetResponse) String() string {
	return string(m)
}

// ReopenLogFilesResponse is message from reopen-log-files command
type ReopenLogFilesResponse client.InformationResponse

// String is a helper method to turn go type into struct
func (r ReopenLogFilesResponse) String() string {
	return string(r)
}

// RulesetReloadRulesResponse is message from ruleset-reload-rules command
type RulesetReloadRulesResponse client.InformationResponse

// String is a helper method to turn go type into struct
func (r RulesetReloadRulesResponse) String() string {
	return string(r)
}

// RulesetReloadTimeResponse is message from ruleset-reload-time (this is usually in a slice in the response)
type RulesetReloadTimeResponse struct {
	ID         int    `json:"id"`
	LastReload string `json:"last_reload"`
}

// RulesetStatsResponse is message from ruleset-stats (this is usually in a slice in the response)
type RulesetStatsResponse struct {
	ID          int `json:"id"`
	RulesFailed int `json:"rules_failed"`
	RulesLoaded int `json:"rules_loaded"`
}

// RulesetFailedRulesResponse is message from ruleset-failed-rules (this is usually in a slice in the response)
type RulesetFailedRulesResponse struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Rule     string `json:"rule"`
	TenantID int    `json:"tenant_id"`
}

// ListHostBitsResponse holds response from list-hostbits <ip>
type ListHostBitsResponse struct {
	Count    int `json:"count"`
	Hostbits []struct {
		Expire int    `json:"expire"`
		Name   string `json:"name"`
	} `json:"hostbits"`
}

// AddOrRemoveHostBitResponse holds response from adding or removing a hostbit <ip> <bitname>
type AddOrRemoveHostBitResponse string

// String is a helper method to turn go type into struct
func (a AddOrRemoveHostBitResponse) String() string {
	return string(a)
}
