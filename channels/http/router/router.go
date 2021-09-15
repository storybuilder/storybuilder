package router

import (
	chi "github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/channels/http/controllers"
	"github.com/storybuilder/storybuilder/channels/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container) *chi.Mux {
	// create new router
	r := chi.NewRouter()

	// initialize middleware
	corsMiddleware := middleware.NewCORSMiddleware()
	requestCheckerMiddleware := middleware.NewRequestCheckerMiddleware(ctr)
	requestAlterMiddleware := middleware.NewRequestAlterMiddleware()
	metricsMiddleware := middleware.NewMetricsMiddleware()

	// add middleware to router
	// NOTE: middleware will execute in the order they are added to the router

	// add metrics middleware first
	r.Use(metricsMiddleware.Middleware)
	r.Use(corsMiddleware)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Logger)
	r.Use(requestCheckerMiddleware.Middleware)
	r.Use(requestAlterMiddleware.Middleware)

	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes

	// api info
	r.Get("/", apiController.GetInfo)

	// sample
	r.Get("/samples", sampleController.Get)
	r.Get("/samples/{id:[0-9]+}", sampleController.GetByID)
	r.Post("/samples", sampleController.Add)
	r.Put("/samples/{id:[0-9]+}", sampleController.Edit)
	r.Delete("/samples/{id:[0-9]+}", sampleController.Delete)

	return r
}
