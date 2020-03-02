package client

import (
	"encoding/json"
)

// DumpCountersCommand performs the dump-counters command
func (s *socket) DumpCountersCommand() (*DumpCountersResponse, error) {
	// create and marshal the "dump-counters" socket message with no args
	response, err := s.DoCommand(dumpCountersCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the reload rules response string
	var counters *DumpCountersResponse
	if err := json.Unmarshal(response.Message, &counters); err != nil {
		return nil, err
	}

	return counters, nil
}
