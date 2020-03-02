package client

import (
	"encoding/json"
)

// VersionCommand gets version of suricata and sets the version in the pointer
func (s *socket) VersionCommand() (string, error) {
	// create and marshal the "version" socket message with no args
	response, err := s.DoCommand(versionCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the version response string
	var version VersionResponse
	if err := json.Unmarshal(response.Message, &version); err != nil {
		return "", err
	}

	return version.String(), nil
}
