package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/cgrs/ecommerce-service-starter/server"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	RootHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := strings.Builder{}
	enc := json.NewEncoder(&expected)

	if err := enc.Encode(map[string]interface{}{"status": 200, "message": "OK"}); err != nil {
		t.Fatal(err)
	}

	if rr.Body.String() != expected.String() {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected.String())
	}
}
