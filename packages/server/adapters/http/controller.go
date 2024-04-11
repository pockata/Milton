package http

import (
	"encoding/json"
	"milton/core/ports"
	"milton/core/services"
	"net/http"
)

type HTTPController struct {
	app services.App
	log ports.Logger
}

type HTTPControllerConfig struct {
	App    services.App
	Logger ports.Logger
}

func NewHTTPController(cfg HTTPControllerConfig) HTTPController {
	return HTTPController{
		app: cfg.App,
		log: cfg.Logger,
	}
}

type ResponseType struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

// handle API responses
func (c HTTPController) Response(rw http.ResponseWriter, r *http.Request, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(data)
}

// handle successful API responses
func (c HTTPController) SuccessResponse(rw http.ResponseWriter, r *http.Request, data interface{}) {
	resp := ResponseType{
		Data:   data,
		Errors: nil,
	}

	c.Response(rw, r, http.StatusOK, resp)
}

// handle erroneous API responses
func (c HTTPController) ErrorResponse(rw http.ResponseWriter, r *http.Request, errs ...error) {
	errStrs := make([]string, 0, len(errs))
	for _, e := range errs {
		msg := e.Error()
		errStrs = append(errStrs, msg)
		c.log.Error(msg, "req-id", rw.Header().Get("X-Request-Id"))
	}

	resp := ResponseType{
		Data:   nil,
		Errors: errStrs,
	}

	c.Response(rw, r, http.StatusBadRequest, resp)
}

func (c HTTPController) ValidParams(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return false
		}
	}

	return true
}
