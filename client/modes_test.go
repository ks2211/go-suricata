package client

import (
	"context"
	"reflect"
	"testing"
)

func TestRunningModeCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := "workers"

	runningMode, err := s.RunningModeCommand(context.TODO())
	if err != nil {
		t.Fatalf("error doing running-mode command %v", err)
	}
	u := reflect.TypeOf(runningMode)
	if u.Name() != "string" {
		t.Fatalf("expected type %T got %v", runningMode, u.Name())
	}
	if runningMode == "" || runningMode != expectedOutput {
		t.Fatalf("expected running-mode output %v got %v", expectedOutput, runningMode)
	}
}

func TestCaptureModeCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := "AF_PACKET_DEV"

	captureMode, err := s.CaptureModeCommand(context.TODO())
	if err != nil {
		t.Fatalf("error doing capture-mode command %v", err)
	}
	u := reflect.TypeOf(captureMode)
	if u.Name() != "string" {
		t.Fatalf("expected type %T got %v", captureMode, u.Name())
	}
	if captureMode == "" || captureMode != expectedOutput {
		t.Fatalf("expected capture-mode output %v got %v", expectedOutput, captureMode)
	}
}
