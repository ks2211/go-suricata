package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	runningModeCommand string = "running-mode"
	captureModeCommand string = "capture-mode"
)

// RunningModeCommand gets running-mode string
func (s *SocketV3) RunningModeCommand() (string, error) {
	// create and marshal the "running-mode" socket message with no args
	response, err := s.DoCommand(runningModeCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the running mode response string
	var runningMode client.RunningModeResponse
	if err := json.Unmarshal(response.Message, &runningMode); err != nil {
		return "", err
	}

	return runningMode.String(), nil
}

// CaptureModeCommand gets capture-mode string
func (s *SocketV3) CaptureModeCommand() (string, error) {
	// create and marshal the "capture-mode" socket message with no args
	response, err := s.DoCommand(captureModeCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the caputre mode response string
	var captureMode client.CaptureModeResponse
	if err := json.Unmarshal(response.Message, &captureMode); err != nil {
		return "", err
	}

	return captureMode.String(), nil
}
