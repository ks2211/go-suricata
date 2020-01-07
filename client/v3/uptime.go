package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	uptimeCommand string = "uptime"
)

// UptimeCommand gets uptime of suricata
func (s *SocketV3) UptimeCommand() (int, error) {
	// create and marshal the "uptime" socket message with no args
	response, err := s.DoCommand(uptimeCommand, nil)
	if err != nil {
		return 0, err
	}
	// unmarshal into the uptime response string
	var uptime client.UptimeResponse
	if err := json.Unmarshal(response.Message, &uptime); err != nil {
		return 0, err
	}

	return int(uptime), nil
}
