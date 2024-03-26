package foundation

import (
	"fmt"
	"net/http"

	"github.com/lucsky/cuid"
)

const reqID = "X-Request-Id"

func SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(reqID)

		// generate a new request id if we don't already have one
		if id == "" {
			id = fmt.Sprintf("rq%s", cuid.New())
		}

		w.Header().Set(reqID, id)

		next.ServeHTTP(w, r)
	})
}
