package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	got, _ := GetAPIKey(req)
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
