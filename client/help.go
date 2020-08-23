package client

import "context"

const (
	help string = "help"
)

// HelpCommand does "help" command which has the same response as "command-list"
func (s *Socket) HelpCommand(ctx context.Context) (CommandListResponse, error) {
	return s.CommandListCommand(ctx)
}
