package sample

import (
	"github.com/storybuilder/storybuilder/app/container"
	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/boundary/repositories"
)

// Sample contains all use cases for samples.
type Sample struct {
	transaction      adapters.DBTxAdapterInterface
	sampleRepository repositories.SampleRepositoryInterface
}

// NewSample creates a new instance of sample use case.
func NewSample(ctr *container.Container) *Sample {
	return &Sample{
		transaction:      ctr.Adapters.DBTxAdapter,
		sampleRepository: ctr.Repositories.SampleRepository,
	}
}
