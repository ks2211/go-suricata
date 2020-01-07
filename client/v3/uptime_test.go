package v3

import (
	"reflect"
	"testing"
)

func TestUptimeCommand(t *testing.T) {
	s, err := NewSocketV3(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	if err := s.ConnectSocket(); err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	uptime, err := s.UptimeCommand()
	if err != nil {
		t.Fatalf("error doing uptime command %v", err)
	}
	u := reflect.TypeOf(uptime)
	if u.Name() != "int" {
		t.Fatalf("expected type %T got %v", uptime, u.Name())
	}
	if uptime == 0 {
		t.Fatalf("uptime should be greater than 0")
	}
}
