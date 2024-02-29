package middleware

import "net/http"

func Register(mux *http.ServeMux) http.Handler {
	return loggingMiddleware(panicMiddleware(mux))
}
