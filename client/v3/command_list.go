package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	commandListCommand string = "command-list"
)

// CommandListCommand gets the list of commands available for this version of suricata and adds it to list of commands
func (s *SocketV3) CommandListCommand() ([]string, error) {
	// create and marshal the "command-list" socket message with no args
	response, err := s.DoCommand(commandListCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the command list struct
	var commandList *client.CommandListResponse
	if err := json.Unmarshal(response.Message, &commandList); err != nil {
		return nil, err
	}
	// update the pointers command slice with the response

	return commandList.Commands, nil
}
