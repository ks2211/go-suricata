package client

import (
	"encoding/json"
)

// ConfGetRequest holds the config key item for conf-get command
type ConfGetRequest struct {
	Variable string `json:"variable"`
}

// ConfGetCommand performs a fetch on a config item using a key string, only returns string/int/bool values
func (s *socket) ConfGetCommand(configItem string) (string, error) {
	// create and marshal the "conf-get" socket message with the config item struct
	response, err := s.DoCommand(confGetCommand, ConfGetRequest{
		configItem,
	})
	if err != nil {
		return "", err
	}
	// unmarshal into the conf get response string
	var configResp ConfGetResponse
	if err := json.Unmarshal(response.Message, &configResp); err != nil {
		return "", err
	}

	return configResp.String(), nil
}
