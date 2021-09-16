package middleware

import (
	"net/http"

	"github.com/storybuilder/storybuilder/app/container"
)

func Init(ctr *container.Container) []func(http.Handler) http.Handler {
	defer println("initialized Middleware")
	metrics := NewMetricsMiddleware()
	cors := NewCORSMiddleware()
	logger := NewLoggerMiddleware()
	requestID := NewRequestIDMiddleware()
	requestChecker := NewRequestCheckerMiddleware(ctr)
	requestAlter := NewRequestAlterMiddleware()

	// NOTE: middleware will execute in the order they are added to the router
	return []func(http.Handler) http.Handler{
		// add metrics middleware first
		metrics.Middleware,
		cors.Middleware,
		logger.Middleware,
		requestID.Middleware,
		requestChecker.Middleware,
		requestAlter.Middleware,
	}
}
