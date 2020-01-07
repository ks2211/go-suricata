package v3

import (
	"encoding/json"
	"github.com/ks2211/go-suricata/client"
)

const (
	confGetCommand = "conf-get"
)

// ConfGetRequest holds the config key item for conf-get command
type ConfGetRequest struct {
	Variable string `json:"variable"`
}

// ConfGetCommand performs a fetch on a config item using a key string, only returns string/int/bool values
func (s *SocketV3) ConfGetCommand(configItem string) (string, error) {
	// create and marshal the "conf-get" socket message with the config item struct
	response, err := s.DoCommand(confGetCommand, ConfGetRequest{
		configItem,
	})
	if err != nil {
		return "", err
	}
	// unmarshal into the reload rules response string
	var configResp client.ConfGetResponse
	if err := json.Unmarshal(response.Message, &configResp); err != nil {
		return "", err
	}

	return configResp.String(), nil
}
