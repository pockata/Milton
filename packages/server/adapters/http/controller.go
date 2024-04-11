package http

import (
	"encoding/json"
	"milton/core/ports"
	"net/http"
)

type HTTPController struct {
	log        ports.Logger
	flowerPots ports.FlowerPotService
	units      ports.UnitService
	jobs       ports.JobService
}

func NewHTTPController(
	log ports.Logger,
	flowerPots ports.FlowerPotService,
	units ports.UnitService,
	jobs ports.JobService,
) HTTPController {
	return HTTPController{
		log:        log,
		flowerPots: flowerPots,
		units:      units,
		jobs:       jobs,
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
