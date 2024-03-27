package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ardanlabs/conf/v3"

	"milton"
	"milton/app"
	"milton/foundation"
	"milton/helpers"
	"milton/routes"
	"milton/storage"
)

func main() {
	// Construct the application logger.
	log := foundation.NewStandardLogger("MILTON-API")

	if err := run(&log); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(log milton.Logger) error {
	cfg := struct {
		conf.Version

		Web struct {
			Host string `conf:"default:0.0.0.0,env:HOST"`
			Port int    `conf:"default:8888,env:PORT"`
		}

		DB struct {
			File string `conf:"default:sqlite.db,env:DB_FILE"`
		}
	}{
		Version: conf.Version{
			Build: milton.Build,
			Desc:  "Milton API server",
		},
	}

	help, err := conf.Parse("", &cfg)

	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		return fmt.Errorf("parsing config: %w", err)
	}

	log.Info("starting service", "version", milton.Build)

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info("startup", "config", out)

	db := storage.NewDB(cfg.DB.File, log)
	dbInstance, err := db.Connect()

	if err != nil {
		return fmt.Errorf("couldn't connect to the database: %w", err)
	}

	ctrl := routes.NewController(routes.ControllerConfig{
		Logger: log,
		App: app.NewApp(app.AppConfig{
			FlowerPotService: storage.NewFlowerPotService(dbInstance),
			UnitService:      storage.NewUnitService(dbInstance),
			JobService:       storage.NewJobService(dbInstance),
		}),
	})

	router := http.NewServeMux()
	w := helpers.CreateAPIWrapHandler(dbInstance)

	router.HandleFunc("GET /query-active-units", w(routes.QueryActiveUnits))

	// units
	router.HandleFunc("GET /get-paired-units", ctrl.GetPairedUnits)
	router.HandleFunc("POST /pair-unit", ctrl.PairUnit)
	router.HandleFunc("POST /unpair-unit", ctrl.UnpairUnit)

	// pots
	router.HandleFunc("POST /add-pot", ctrl.AddPot)
	router.HandleFunc("GET /get-pots/{UnitID}", ctrl.GetPots)
	router.HandleFunc("POST /rename-pot", ctrl.RenamePot)
	router.HandleFunc("POST /remove-pot", ctrl.RemovePot)

	// watering jobs
	router.HandleFunc("POST /add-job", ctrl.AddJob)
	router.HandleFunc("POST /remove-job", ctrl.RemoveJob)
	router.HandleFunc("POST /update-job", ctrl.UpdateJob)
	router.HandleFunc("GET /get-jobs", ctrl.GetJobs)
	router.HandleFunc("GET /get-job/{JobID}", ctrl.GetJob)

	wrappedRouter := foundation.SetRequestID(router)
	wrappedRouter = foundation.RequestLogger(wrappedRouter, log)

	addr := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)
	return http.ListenAndServe(addr, wrappedRouter)
}

// func main2() {
// 	var m mqtt.MQTT
// 	var db models.DB
//
// 	Config := config.Read()
//
// 	db.Setup()
// 	m.Setup(Config.MQTT)
//
// 	router := mux.NewRouter().StrictSlash(true)
//
// 	api := router.PathPrefix("/api/").Subrouter()
//
// 	api.Use(helpers.CORSHeaders(api, Config.CORS))
//
// 	w := helpers.CreateAPIWrapHandler(db)
//
// 	api.HandleFunc("/query-active-units", w(routes.QueryActiveUnits)).Methods(http.MethodGet)
//
// 	// units
// 	api.HandleFunc("/get-paired-units", w(routes.GetPairedUnits)).Methods(http.MethodGet)
// 	api.HandleFunc("/pair-unit", w(routes.PairUnit)).Methods(http.MethodPost)
// 	api.HandleFunc("/unpair-unit", w(routes.UnpairUnit)).Methods(http.MethodPost)
//
// 	// pots
// 	api.HandleFunc("/add-pot", w(routes.AddPot)).Methods(http.MethodPost)
// 	api.HandleFunc("/get-pots/{UnitID}", w(routes.GetPots)).Methods(http.MethodGet)
// 	api.HandleFunc("/update-pot", w(routes.UpdatePot)).Methods(http.MethodPost)
// 	api.HandleFunc("/remove-pot", w(routes.RemovePot)).Methods(http.MethodPost)
//
// 	// watering jobs
// 	api.HandleFunc("/add-job", w(routes.AddJob)).Methods(http.MethodPost)
// 	api.HandleFunc("/remove-job", w(routes.RemoveJob)).Methods(http.MethodPost)
// 	api.HandleFunc("/update-job", w(routes.UpdateJob)).Methods(http.MethodPost)
// 	api.HandleFunc("/get-jobs", w(routes.GetJobs)).Methods(http.MethodGet)
// 	api.HandleFunc("/get-job/{JobID}", w(routes.GetJob)).Methods(http.MethodGet)
//
// 	// TODO: Support PORT env variable
// 	// httpPort := os.Getenv("PORT")
// 	// 	if httpPort == "" {
// 	// 		httpPort = "8080"
// 	// 	}
// 	log.Fatal(http.ListenAndServe(Config.Server.Address, router))
// }
