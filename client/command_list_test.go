package client

import (
	"context"
	"testing"
)

func TestCommandListCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
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
		"ruleset-reload-rules",
		"ruleset-reload-nonblocking",
		"ruleset-reload-time",
		"ruleset-stats",
		"ruleset-failed-rules",
		"register-tenant-handler",
		"unregister-tenant-handler",
		"register-tenant",
		"reload-tenant",
		"unregister-tenant",
		"add-hostbit",
		"remove-hostbit",
		"list-hostbit",
		"reopen-log-files",
		"memcap-set",
		"memcap-show",
		"memcap-list",
		"dataset-add",
		"iface-stat",
		"iface-list",
		"iface-bypassed-stat",
		"ebpf-bypassed-stat",
	}

	commands, err := s.CommandListCommand(context.TODO())
	if err != nil {
		t.Fatalf("error doing command-list command %v", err)
	}
	if len(commands.Commands) == 0 || len(commands.Commands) != len(expectedOutput) {
		t.Fatalf("error command list is empty or does not match expected, expected %v, got %v", len(expectedOutput), len(commands.Commands))
	}

}
