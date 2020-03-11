package client

const (
	help string = "help"
)

// HelpCommand does "help" command which has the same response as "command-list"
func (s *Socket) HelpCommand() (*CommandListResponse, error) {
	return s.CommandListCommand()
}
