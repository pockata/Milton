package main

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
func apiResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}

// handle successful API responses
func successResponse(rw http.ResponseWriter, data interface{}) {
	resp := SuccessResponseType{
		Error: false,
		Data:  data,
	}

	apiResponse(rw, http.StatusOK, resp)
}

// handle erroneous API responses
func errorResponse(rw http.ResponseWriter, error error) {
	resp := ErrorResponseType{
		Error: error,
		Msg:   fmt.Sprintf("%v", error),
	}

	apiResponse(rw, http.StatusBadRequest, resp)
}
