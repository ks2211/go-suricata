package client

import (
	"context"
	"reflect"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedVersion := "5.0.1 RELEASE"

	version, err := s.VersionCommand(context.TODO())
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
