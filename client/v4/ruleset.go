package v4

import (
	"encoding/json"
)

// Constants
const (
	rulesetReoladRulesCommand       string = "ruleset-reload-rules"
	rulesetReloadNonBlockingCommand string = "ruleset-reload-nonblocking"
	rulesetReloadTimeCommand        string = "ruleset-reload-time"
	rulesetStatsCommand             string = "ruleset-stats"
	rulesetFailedRulesCommand       string = "ruleset-failed-rules"
)

// RulesetReloadRulesCommand does ruleset reload
func (s *SocketV4) RulesetReloadRulesCommand() (string, error) {
	// create and marshal the "ruleset-reload-rules" socket message with
	response, err := s.DoCommand(rulesetReoladRulesCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var rulesetReloadRules RulesetReloadRulesResponse
	if err := json.Unmarshal(response.Message, &rulesetReloadRules); err != nil {
		return "", err
	}
	return rulesetReloadRules.String(), nil
}

// RulesetReloadNonBlockingCommand does ruleset reload without blocking
func (s *SocketV4) RulesetReloadNonBlockingCommand() (string, error) {
	// create and marshal the "ruleset-reload-nonblocking" socket message with
	response, err := s.DoCommand(rulesetReloadNonBlockingCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var rulesetReloadRules RulesetReloadRulesResponse
	if err := json.Unmarshal(response.Message, &rulesetReloadRules); err != nil {
		return "", err
	}
	return rulesetReloadRules.String(), nil
}

// RulesetReloadTimeCommand gets reload time of ruleset reload
func (s *SocketV4) RulesetReloadTimeCommand() ([]RulesetReloadTimeResponse, error) {
	// create and marshal the "ruleset-reload-time" socket message with
	response, err := s.DoCommand(rulesetReloadTimeCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface stat struct
	var rulesetReloadTime []RulesetReloadTimeResponse
	if err := json.Unmarshal(response.Message, &rulesetReloadTime); err != nil {
		return nil, err
	}
	return rulesetReloadTime, nil
}

// RulesetStatsCommand gets ruleset stats
func (s *SocketV4) RulesetStatsCommand() ([]RulesetStatsResponse, error) {
	// create and marshal the "ruleset-stats" socket message with
	response, err := s.DoCommand(rulesetStatsCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface stat struct
	var rulesetStats []RulesetStatsResponse
	if err := json.Unmarshal(response.Message, &rulesetStats); err != nil {
		return nil, err
	}
	return rulesetStats, nil
}

// RulesetFailedRulesCommand does ruleset failed rules
func (s *SocketV4) RulesetFailedRulesCommand() ([]RulesetFailedRulesResponse, error) {
	// create and marshal the "ruleset-failed-rules" socket message with
	response, err := s.DoCommand(rulesetFailedRulesCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the slice of ruleset failed rules struct
	var rulesetFailed []RulesetFailedRulesResponse
	if err := json.Unmarshal(response.Message, &rulesetFailed); err != nil {
		return nil, err
	}
	return rulesetFailed, nil
}
