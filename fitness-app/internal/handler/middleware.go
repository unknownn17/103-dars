
package handler

import (
	"net/http"

	"fitness/pkg/logger"
)

func LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Method + " " + r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
