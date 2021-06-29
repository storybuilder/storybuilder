package container

import (
	"fmt"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/externals/services"
)

var resolvedServices Services

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {
	resolveSampleService(getConfigByName(cfgs, "sample"))

	return resolvedServices
}

// resolveSampleService resolves the sample service.
func resolveSampleService(cfg config.ServiceConfig) {
	resolvedServices.SampleService = services.NewSampleService(cfg)
}

// getConfigByName returns the service config by name of the service.
func getConfigByName(cfgs []config.ServiceConfig, name string) config.ServiceConfig {
	for i := range cfgs {
		if cfgs[i].Name == name {
			return cfgs[i]
		}
	}

	// must panic if the config is not found
	panic(fmt.Sprintf("Cannot find service configurations for `%s` service", name))
}
