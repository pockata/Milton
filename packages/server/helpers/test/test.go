package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
)

func executeRequest(
	router mux.Router,
	req *http.Request,
) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}
