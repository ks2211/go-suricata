package client

import "context"

const (
	reopenLogFiles string = "reopen-log-files"
)

// ReopenLogFilesResponse is message from reopen-log-files command
type ReopenLogFilesResponse StringResponse

// String is a helper method to turn go type into struct
func (r ReopenLogFilesResponse) String() string {
	return string(r)
}

// ReopenLogFilesCommand reopens log files via the reopen-logfiles command
// Not implemented in v3
func (s *Socket) ReopenLogFilesCommand(ctx context.Context) (string, error) {
	reopenLogFilesResp := new(ReopenLogFilesResponse)
	err := s.DoCommand(ctx, reopenLogFiles, nil, reopenLogFilesResp)
	return reopenLogFilesResp.String(), err
}
