package http

import (
	"net/http"
)

func AddRoutes(api *http.ServeMux, ctrl HTTPController) {
	// w := helpers.CreateAPIWrapHandler(dbInstance)

	// api.HandleFunc("GET /query-active-units", w(routes.QueryActiveUnits))

	// units
	api.HandleFunc("GET /get-paired-units", ctrl.GetPairedUnits)
	api.HandleFunc("POST /pair-unit", ctrl.PairUnit)
	api.HandleFunc("POST /unpair-unit", ctrl.UnpairUnit)

	// pots
	api.HandleFunc("POST /add-pot", ctrl.AddPot)
	api.HandleFunc("GET /get-pots/{UnitID}", ctrl.GetPots)
	api.HandleFunc("POST /rename-pot", ctrl.RenamePot)
	api.HandleFunc("POST /remove-pot", ctrl.RemovePot)

	// watering jobs
	api.HandleFunc("POST /add-job", ctrl.AddJob)
	api.HandleFunc("POST /remove-job", ctrl.RemoveJob)
	api.HandleFunc("POST /update-job", ctrl.UpdateJob)
	api.HandleFunc("GET /get-jobs", ctrl.GetJobs)
	api.HandleFunc("GET /get-job/{JobID}", ctrl.GetJob)
}
