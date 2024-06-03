package internalhttp

import (
	"net/http"
	"strings"
	"time"
)

func loggingMiddleware(next http.Handler, logg Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		logg.Info(strings.Join(
			[]string{
				r.RemoteAddr,
				time.Now().String(),
				r.Method,
				r.URL.Path,
				r.Proto,
				time.Since(start).String(),
				r.UserAgent(),
			},
			" ",
		))

		next.ServeHTTP(w, r)
	})
}
