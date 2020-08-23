package client

import "context"

const (
	commandList string = "command-list"
)

// CommandListResponse is the response from "command-list"
type CommandListResponse struct {
	Commands []string `json:"commands"`
	Count    int      `json:"count"`
}

// CommandListCommand gets the list of commands available for this version of suricata and adds it to list of commands
func (s *Socket) CommandListCommand(ctx context.Context) (CommandListResponse, error) {
	commandListResp := CommandListResponse{}
	err := s.DoCommand(ctx, commandList, nil, &commandListResp)
	return commandListResp, err
}
