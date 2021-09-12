package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewLoggerMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}
