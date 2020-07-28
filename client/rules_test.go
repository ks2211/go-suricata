package client

import (
	"context"
	"reflect"
	"testing"
)

func TestReloadRulesCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := "done"

	reload, err := s.ReloadRulesCommand(context.TODO())
	if err != nil {
		t.Fatalf("error doing reload-rules command %v", err)
	}
	u := reflect.TypeOf(reload)
	if u.Name() != "string" {
		t.Fatalf("expected type %T got %v", reload, u.Name())
	}
	if reload == "" || reload != expectedOutput {
		t.Fatalf("expected reload-rules output %v got %v", expectedOutput, reload)
	}
}
