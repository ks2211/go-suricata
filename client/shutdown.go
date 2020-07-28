package client

import "context"

const (
	shutdown string = "shutdown"
)

// ShutdownResponse is response from "shutdown"
type ShutdownResponse StringResponse

// String is a helper method to convert a go type into string
func (s ShutdownResponse) String() string {
	return string(s)
}

// ShutdownCommand performs a shutdown of suricata and close the net conn
func (s *Socket) ShutdownCommand(ctx context.Context) (string, error) {
	shutdownResp := new(ShutdownResponse)
	err := s.DoCommand(ctx, shutdown, nil, shutdownResp)
	defer s.Close()
	return shutdownResp.String(), err
}
