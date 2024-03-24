package routes

import (
	"milton/app"
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
