package client

import "context"

const (
	rulesetReoladRules       string = "ruleset-reload-rules"
	rulesetReloadNonBlocking string = "ruleset-reload-nonblocking"
	rulesetReloadTime        string = "ruleset-reload-time"
	rulesetStats             string = "ruleset-stats"
	rulesetFailedRules       string = "ruleset-failed-rules"
)

// RulesetReloadRulesResponse is message from ruleset-reload-rules command
type RulesetReloadRulesResponse StringResponse

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

// RulesetReloadRulesCommand does ruleset reload
// Not implemented in v3
func (s *Socket) RulesetReloadRulesCommand(ctx context.Context) (string, error) {
	rulesetReloadRulesResp := new(RulesetReloadRulesResponse)
	err := s.DoCommand(ctx, rulesetReoladRules, nil, rulesetReloadRulesResp)
	return rulesetReloadRulesResp.String(), err
}

// RulesetReloadNonBlockingCommand does ruleset reload without blocking
// Not implemented in v3
func (s *Socket) RulesetReloadNonBlockingCommand(ctx context.Context) (string, error) {
	rulesetReloadRulesResp := new(RulesetReloadRulesResponse)
	err := s.DoCommand(ctx, rulesetReloadNonBlocking, nil, rulesetReloadRulesResp)
	return rulesetReloadRulesResp.String(), err
}

// RulesetReloadTimeCommand gets reload time of ruleset reload
// Not implemented in v3
func (s *Socket) RulesetReloadTimeCommand(ctx context.Context) ([]RulesetReloadTimeResponse, error) {
	rulesetReloadTimeResp := []RulesetReloadTimeResponse{}
	err := s.DoCommand(ctx, rulesetReloadTime, nil, &rulesetReloadTimeResp)
	return rulesetReloadTimeResp, err
}

// RulesetStatsCommand gets ruleset stats
// Not implemented in v3
func (s *Socket) RulesetStatsCommand(ctx context.Context) ([]RulesetStatsResponse, error) {
	rulesetStatsResp := []RulesetStatsResponse{}
	err := s.DoCommand(ctx, rulesetStats, nil, &rulesetStatsResp)
	return rulesetStatsResp, err
}

// RulesetFailedRulesCommand does ruleset failed rules
// Not implemented in v3
func (s *Socket) RulesetFailedRulesCommand(ctx context.Context) ([]RulesetFailedRulesResponse, error) {
	rulesetFailedResp := []RulesetFailedRulesResponse{}
	err := s.DoCommand(ctx, rulesetFailedRules, nil, &rulesetFailedResp)
	return rulesetFailedResp, err
}
