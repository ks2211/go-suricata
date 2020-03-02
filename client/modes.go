package client

import (
	"encoding/json"
)

// RunningModeCommand gets running-mode string
func (s *socket) RunningModeCommand() (string, error) {
	// create and marshal the "running-mode" socket message with no args
	response, err := s.DoCommand(runningModeCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the running mode response string
	var runningMode RunningModeResponse
	if err := json.Unmarshal(response.Message, &runningMode); err != nil {
		return "", err
	}

	return runningMode.String(), nil
}

// CaptureModeCommand gets capture-mode string
func (s *socket) CaptureModeCommand() (string, error) {
	// create and marshal the "capture-mode" socket message with no args
	response, err := s.DoCommand(captureModeCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the caputre mode response string
	var captureMode CaptureModeResponse
	if err := json.Unmarshal(response.Message, &captureMode); err != nil {
		return "", err
	}

	return captureMode.String(), nil
}
