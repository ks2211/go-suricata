package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	dumpCountersCommand string = "dump-counters"
)

// DumpCountersCommand performs the dump-counters command
func (s *SocketV3) DumpCountersCommand() (client.DumpCountersResponse, error) {
	// create and marshal the "dump-counters" socket message with no args
	response, err := s.DoCommand(dumpCountersCommand, nil)
	if err != nil {
		return client.DumpCountersResponse{}, err
	}
	// unmarshal into the reload rules response string
	var counters client.DumpCountersResponse
	if err := json.Unmarshal(response.Message, &counters); err != nil {
		return client.DumpCountersResponse{}, err
	}

	return counters, nil
}
