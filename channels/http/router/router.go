package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/channels/http/controllers"
	"github.com/storybuilder/storybuilder/channels/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container) *mux.Router {
	// create new router
	r := mux.NewRouter()

	// initialize middleware
	loggerMiddleware := middleware.NewLoggerMiddleware()
	requestCheckerMiddleware := middleware.NewRequestCheckerMiddleware(ctr)
	requestAlterMiddleware := middleware.NewRequestAlterMiddleware()
	metricsMiddleware := middleware.NewMetricsMiddleware()

	// add middleware to router
	// NOTE: middleware will execute in the order they are added to the router

	// add metrics middleware first
	r.Use(metricsMiddleware.Middleware)

	// add CORS middleware
	r.Use(mux.CORSMethodMiddleware(r))

	r.Use(loggerMiddleware.Middleware)
	r.Use(requestCheckerMiddleware.Middleware)
	r.Use(requestAlterMiddleware.Middleware)

	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes

	// api info
	r.HandleFunc("/", apiController.GetInfo).Methods(http.MethodGet)

	// sample
	r.HandleFunc("/samples", sampleController.Get).Methods(http.MethodGet)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/samples", sampleController.Add).Methods(http.MethodPost)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Edit).Methods(http.MethodPut)
	r.HandleFunc("/samples/{id:[0-9]+}", sampleController.Delete).Methods(http.MethodDelete)

	return r
}
