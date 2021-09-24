package middlewares

import "net/http"

type JsonMiddleware struct {
	handler http.Handler
}

func (j *JsonMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	j.handler.ServeHTTP(rw, r)
}

func WithJson(next http.Handler) *JsonMiddleware {
	return &JsonMiddleware{next}
}
