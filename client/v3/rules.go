package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	reloadRulesCommand string = "reload-rules"
)

// ReloadRulesCommand performs a reload of rules without restarting suricata
func (s *SocketV3) ReloadRulesCommand() (string, error) {
	// create and marshal the "reload-rules" socket message with no args
	response, err := s.DoCommand(reloadRulesCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the reload rules response string
	var reloadRules client.ReloadRulesResponse
	if err := json.Unmarshal(response.Message, &reloadRules); err != nil {
		return "", err
	}

	return reloadRules.String(), nil
}
