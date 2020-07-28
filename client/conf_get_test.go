package client

import (
	"context"
	"testing"
)

func TestConfGetCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	cases := []struct {
		confItem string
		output   string
	}{
		{
			"unix-command.filename",
			"/var/run/suricata-command.socket",
		},
		{
			"vars.port-groups.HTTP_PORTS",
			"80",
		},
		{
			"vars.port-groups.SSH_PORTS",
			"22",
		},
		{
			"stats.enabled",
			"yes",
		},
	}

	for _, tt := range cases {
		confItem, err := s.ConfGetCommand(context.TODO(), ConfGetRequest{
			tt.confItem,
		})
		if err != nil {
			t.Fatalf("error doing conf-get command for %v, error %v", tt.confItem, err)
		}
		if confItem != tt.output {
			t.Fatalf("expected conf-get %v output %v got %v", tt.confItem, tt.output, confItem)
		}
	}
}
