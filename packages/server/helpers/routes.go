package helpers

import (
	"database/sql"
	"net/http"
)

type APIWrapperFunc func(http.ResponseWriter, *http.Request, *sql.DB)

func CreateAPIWrapHandler(db *sql.DB) func(APIWrapperFunc) http.HandlerFunc {
	return func(handler APIWrapperFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			handler(rw, r, db)
		}
	}
}
