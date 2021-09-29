package middlewares

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	handler http.Handler
	logger  *log.Logger
}

func (l *LoggerMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	t := time.Now()
	srw := NewStatusResponseWriter(rw)
	l.handler.ServeHTTP(srw, r)
	l.logger.Printf("%s %s %d - %v", r.Method, r.URL, srw.statusCode, time.Since(t))
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{w, http.StatusOK}
}

func (srw *statusResponseWriter) WriteHeader(status int) {
	srw.statusCode = status
	srw.ResponseWriter.WriteHeader(status)
}



func WithLogger(next http.Handler, logger *log.Logger) *LoggerMiddleware {
	if logger == nil {
		logger = log.Default()
	}
	return &LoggerMiddleware{next, logger}
}
