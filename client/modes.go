package client

import "context"

const (
	captureMode string = "capture-mode"
	runningMode string = "running-mode"
)

// RunningModeResponse is response from "running-mode"
type RunningModeResponse StringResponse

// String is a helper method to convert a go type into string
func (r RunningModeResponse) String() string {
	return string(r)
}

// CaptureModeResponse is response from "capture-mode"
type CaptureModeResponse StringResponse

// String is a helper method to convert a go type into string
func (c CaptureModeResponse) String() string {
	return string(c)
}

// RunningModeCommand gets running-mode string
func (s *Socket) RunningModeCommand(ctx context.Context) (string, error) {
	runningModeResp := new(RunningModeResponse)
	err := s.DoCommand(ctx, runningMode, nil, runningModeResp)
	return runningModeResp.String(), err
}

// CaptureModeCommand gets capture-mode string
func (s *Socket) CaptureModeCommand(ctx context.Context) (string, error) {
	captureModeResp := new(CaptureModeResponse)
	err := s.DoCommand(ctx, captureMode, nil, captureModeResp)
	return captureModeResp.String(), err
}
