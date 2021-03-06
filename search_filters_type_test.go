package elastic

import (
	"encoding/json"
	"testing"
)

func TestTypeFilter(t *testing.T) {
	f := NewTypeFilter("my_type")
	data, err := json.Marshal(f.Source())
	if err != nil {
		t.Fatalf("marshaling to JSON failed: %v", err)
	}
	got := string(data)
	expected := `{"type":{"value":"my_type"}}`
	if got != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, got)
	}
}
