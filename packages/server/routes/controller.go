package routes

import (
	"encoding/json"
	"milton/app"
	"net/http"
)

type Controller struct {
	app app.App
}

type ControllerConfig struct {
	App app.App
}

func NewController(cfg ControllerConfig) Controller {
	return Controller{
		app: cfg.App,
	}
}

type APIResponseType struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

// handle API responses
func (c Controller) ApiResponse(rw http.ResponseWriter, r *http.Request, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}

// handle successful API responses
func (c Controller) SuccessResponse(rw http.ResponseWriter, r *http.Request, data interface{}) {
	resp := APIResponseType{
		Data:   data,
		Errors: nil,
	}

	c.ApiResponse(rw, r, http.StatusOK, resp)
}

// handle erroneous API responses
func (c Controller) ErrorResponse(rw http.ResponseWriter, r *http.Request, errs ...error) {
	errStrs := make([]string, 0, len(errs))
	for _, e := range errs {
		errStrs = append(errStrs, e.Error())
	}

	resp := APIResponseType{
		Data:   nil,
		Errors: errStrs,
	}

	c.ApiResponse(rw, r, http.StatusBadRequest, resp)
}
