package app

import (
	"devops-study-golang/pkg"
	"devops-study-golang/structure/features/health"
)

type App struct {
	HealthHandler  *health.Handler
}

func New(cfg *pkg.Config) (*App, error) {
	healthHandler := health.NewHandler()

	app := &App{
		HealthHandler: healthHandler,
	}

	return app, nil
}

func (a *App) Close() {
}
