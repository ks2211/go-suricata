package client

import (
	"encoding/json"
)

// ShutdownCommand performs a shutdown of suricata and close the net conn
func (s *socket) ShutdownCommand() (string, error) {
	// create and marshal the "shutdown" socket message with no args
	response, err := s.DoCommand(shutdownCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the shutdown response string
	var shutdown ShutdownResponse
	if err := json.Unmarshal(response.Message, &shutdown); err != nil {
		return "", err
	}
	// close the socket connection
	defer s.Close()

	return shutdown.String(), nil
}
