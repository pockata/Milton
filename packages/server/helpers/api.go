package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResponseType struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

type ErrorResponseType struct {
	Error error  `json:"error"`
	Msg   string `json:"msg"`
}

// handle API responses
func ApiResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}

// handle successful API responses
func SuccessResponse(rw http.ResponseWriter, data interface{}) {
	resp := SuccessResponseType{
		Error: false,
		Data:  data,
	}

	ApiResponse(rw, http.StatusOK, resp)
}

// handle erroneous API responses
func ErrorResponse(rw http.ResponseWriter, error error) {
	resp := ErrorResponseType{
		Error: error,
		Msg:   fmt.Sprintf("%v", error),
	}

	ApiResponse(rw, http.StatusBadRequest, resp)
}

func CheckParams(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}

	return true
}