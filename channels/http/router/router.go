package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/channels/http/controllers"
	"github.com/storybuilder/storybuilder/channels/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container) *chi.Mux {
	// create new router
	r := chi.NewRouter()

	// initialize middleware
	metrics := middleware.NewMetricsMiddleware()
	cors := middleware.NewCORSMiddleware()
	logger := middleware.NewLoggerMiddleware()
	requestID := middleware.NewRequestIDMiddleware()
	requestChecker := middleware.NewRequestCheckerMiddleware(ctr)
	requestAlter := middleware.NewRequestAlterMiddleware()

	// add middleware to router
	// NOTE: middleware will execute in the order they are added to the router

	// add metrics middleware first
	r.Use(metrics.Middleware)
	r.Use(cors.Middleware)
	r.Use(logger.Middleware)
	r.Use(requestID.Middleware)
	r.Use(requestChecker.Middleware)
	r.Use(requestAlter.Middleware)

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
