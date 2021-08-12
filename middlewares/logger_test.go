package middlewares_test

import (
	"log"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	. "github.com/cgrs/ecommerce-service-starter/middlewares"
)

func TestLoggerMiddleware(t *testing.T) {
	rr := httptest.NewRecorder()
	sb := &strings.Builder{}
	m := WithLogger(http.DefaultServeMux, log.New(sb, "", 0))
	m.ServeHTTP(rr, httptest.NewRequest("GET", "/", nil))
	if !strings.Contains(sb.String(), "GET /") {
		t.Errorf("logger received unknown data: got %v want %v", sb.String(), "GET /")
	}
}
