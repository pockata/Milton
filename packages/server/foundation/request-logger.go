package foundation

import (
	"milton"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler, log milton.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := newWrappedResponse(w)

		next.ServeHTTP(&wrapped, r)

		log.Info("Request",
			"status", wrapped.statusCode,
			"method", r.Method,
			"url", getUrl(r),
			"req-id", wrapped.Header().Get("X-Request-Id"),
			"dur", time.Since(start),
		)
	})
}

type wrappedResponse struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponse) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.statusCode = status
}

func newWrappedResponse(w http.ResponseWriter) wrappedResponse {
	return wrappedResponse{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func getUrl(r *http.Request) string {
	qs := ""
	if len(r.URL.RawQuery) > 0 {
		qs = "?" + r.URL.RawQuery
	}

	return r.URL.Path + qs
}
