package v3

import (
	"testing"
)

func TestCommandListCommand(t *testing.T) {
	s, err := NewSocketV3(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	if err := s.ConnectSocket(); err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := []string{
		"shutdown",
		"command-list",
		"help",
		"version",
		"uptime",
		"running-mode",
		"capture-mode",
		"conf-get",
		"dump-counters",
		"reload-rules",
		"register-tenant-handler",
		"unregister-tenant-handler",
		"register-tenant",
		"reload-tenant",
		"unregister-tenant",
		"iface-stat",
		"iface-list",
	}

	commands, err := s.CommandListCommand()
	if err != nil {
		t.Fatalf("error doing command-list command %v", err)
	}
	if len(commands) == 0 || len(commands) != len(expectedOutput) {
		t.Fatalf("error command list is empty or does not match expected, expected %v, got %v", len(expectedOutput), len(commands))
	}

}
