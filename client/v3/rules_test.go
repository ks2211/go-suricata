package v3

import (
	"reflect"
	"testing"
)

func TestReloadRulesCommand(t *testing.T) {
	s, err := NewSocketV3(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	if err := s.ConnectSocket(); err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := "done"

	reload, err := s.ReloadRulesCommand()
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
