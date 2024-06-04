package middleware

import (
	"fmt"
	"net/http"
	log "ozon/pkg"
)

func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l, err := log.GetFromContext(r.Context())
		if err != nil {
			panic(fmt.Errorf("logger object can't be found in the context: %w", err))
		}
		l.Infof("%s request from %s to %s", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
