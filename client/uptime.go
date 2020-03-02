package client

import (
	"encoding/json"
)

// UptimeCommand gets uptime of suricata
func (s *socket) UptimeCommand() (int, error) {
	// create and marshal the "uptime" socket message with no args
	response, err := s.DoCommand(uptimeCommand, nil)
	if err != nil {
		return 0, err
	}
	// unmarshal into the uptime response string
	var uptime UptimeResponse
	if err := json.Unmarshal(response.Message, &uptime); err != nil {
		return 0, err
	}

	return int(uptime), nil
}
