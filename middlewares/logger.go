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
	l.handler.ServeHTTP(rw, r)
	l.logger.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(t))
}

func WithLogger(next http.Handler, logger *log.Logger) *LoggerMiddleware {
	if logger == nil {
		logger = log.Default()
	}
	return &LoggerMiddleware{next, logger}
}
