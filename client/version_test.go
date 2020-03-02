package client

import (
	"reflect"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedVersion := "3.2 RELEASE"

	version, err := s.VersionCommand()
	if err != nil {
		t.Fatalf("error doing version command %v", err)
	}
	u := reflect.TypeOf(version)
	if u.Name() != "string" {
		t.Fatalf("expected type %T got %v", version, u.Name())
	}
	if version == "" || version != expectedVersion {
		t.Fatalf("expected version %v got %v", expectedVersion, version)
	}
}
