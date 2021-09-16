package middleware

import (
	"net/http"

	"github.com/storybuilder/storybuilder/app/container"
)

type Middleware interface {
	Middleware(next http.Handler) http.Handler
}

func Init(ctr *container.Container) (middlewares []func(http.Handler) http.Handler) {
	// NOTE: middleware will execute in the order they are added to the router
	ms := []Middleware{
		// add metrics middleware first
		NewMetricsMiddleware(),
		NewCORSMiddleware(),
		NewLoggerMiddleware(),
		NewRequestIDMiddleware(),
		NewRequestCheckerMiddleware(ctr),
		NewRequestAlterMiddleware(),
	}
	for _, middleware := range ms {
		middlewares = append(middlewares, middleware.Middleware)
	}
	return
}
