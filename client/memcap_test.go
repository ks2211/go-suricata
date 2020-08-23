package client

import (
	"context"
	"testing"
)

func TestMemCapListCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	memcapList, err := s.MemCapListCommand(context.TODO())
	if err != nil {
		t.Fatalf("error doing iface-stat command %v", err)
	}
	if len(memcapList) == 0 {
		t.Fatalf("error want %d got no memcap in list", len(memcapList))
	}
}
