package server_test

import (
	"net/http"
	"testing"

	. "github.com/cgrs/ecommerce-service-starter/server"
)

func TestCreateServer(t *testing.T) {
	s := CreateServer("", http.DefaultServeMux)
	expectedAddr := "localhost:3000"
	if s.Addr != expectedAddr {
		t.Errorf("unknown address, got %v want %v", s.Addr, expectedAddr)
	}
}
