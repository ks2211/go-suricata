package client

import (
	"encoding/json"
)

// ReloadRulesCommand performs a reload of rules without restarting suricata
func (s *socket) ReloadRulesCommand() (string, error) {
	// create and marshal the "reload-rules" socket message with no args
	response, err := s.DoCommand(reloadRulesCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the reload rules response string
	var reloadRules ReloadRulesResponse
	if err := json.Unmarshal(response.Message, &reloadRules); err != nil {
		return "", err
	}

	return reloadRules.String(), nil
}
