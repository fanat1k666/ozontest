package middleware

import (
	"net/http"
	log "ozon/pkg"
)

func AddLogger(next http.HandlerFunc, l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(log.AddToContext(r.Context(), l))
		next.ServeHTTP(w, r)
	}
}
