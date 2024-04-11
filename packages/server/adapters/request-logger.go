package adapters

import (
	"milton/core/ports"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler, log ports.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Info("Request",
			"method", r.Method,
			"url", getUrl(r),
			"req-id", w.Header().Get("X-Request-Id"),
			"dur", time.Since(start),
		)
	})
}

func getUrl(r *http.Request) string {
	qs := ""
	if len(r.URL.RawQuery) > 0 {
		qs = "?" + r.URL.RawQuery
	}

	return r.URL.Path + qs
}
