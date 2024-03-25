package helpers

import (
	"encoding/json"
	"net/http"
)

type APIConfig struct {
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
}

type APIResponseType struct {
	Errors []error     `json:"errors"`
	Data   interface{} `json:"data"`
}

// handle API responses
func ApiResponse(rw http.ResponseWriter, r *http.Request, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}

// handle successful API responses
func SuccessResponse(rw http.ResponseWriter, r *http.Request, data interface{}) {
	resp := APIResponseType{
		Data: data,
	}

	ApiResponse(rw, r, http.StatusOK, resp)
}

// handle erroneous API responses
func ErrorResponse(rw http.ResponseWriter, r *http.Request, errs ...error) {
	resp := APIResponseType{
		Errors: errs,
	}

	ApiResponse(rw, r, http.StatusBadRequest, resp)
}

func ValidParams(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}

	return true
}
