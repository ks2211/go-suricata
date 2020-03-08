package client

import (
	"testing"
)

func TestCreateSocket(t *testing.T) {
	cases := []struct {
		path         string
		expectedPath string
	}{
		{
			"/var/run/suricata-command.socket",
			"/var/run/suricata-command.socket",
		},
	}
	for _, tt := range cases {
		sock, err := CreateSocket(tt.path)
		defer sock.Close()
		if err != nil {
			t.Fatalf("error creating socket client %v", err)
		}
		if sock.Path() != tt.expectedPath {
			t.Fatalf("expected path %v got %v", tt.expectedPath, sock.Path())
		}
	}
}

func TestDoCommand(t *testing.T) {
	cases := []struct {
		cmdName string
		output  string
	}{
		{
			"running-mode",
			"workers",
		},
		{
			"capture-mode",
			"AF_PACKET_DEV",
		},
		{
			"version",
			"3.2 RELEASE",
		},
	}
	s, err := CreateSocket("/var/run/suricata-command.socket")
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()
	for _, tt := range cases {
		respMessage := new(StringResponse)
		err := s.DoCommand(tt.cmdName, nil, respMessage)
		if err != nil {
			t.Fatalf("error doing command %v, error %v", tt.cmdName, err)
		}
		if respMessage.String() != tt.output {
			t.Fatalf("expected %v, got %v", tt.output, respMessage.String())
		}
	}
}
