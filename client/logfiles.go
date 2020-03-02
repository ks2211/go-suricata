package client

import "encoding/json"

// ReopenLogFilesCommand reopens log files via the reopen-logfiles command
// Not implemented in v3
func (s *socket) ReopenLogFilesCommand() (string, error) {
	// create and marshal the "reopen-log-files" socket message with
	response, err := s.DoCommand(reopenLogFilesCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var reopenLogFiles ReopenLogFilesResponse
	if err := json.Unmarshal(response.Message, &reopenLogFiles); err != nil {
		return "", err
	}
	return reopenLogFiles.String(), nil
}
