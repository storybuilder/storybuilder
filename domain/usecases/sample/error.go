package sample

import (
	"fmt"

	err "github.com/storybuilder/storybuilder/domain/errors"
)

func (s *Sample) errorNoSample(id int) error {
	return err.NewDomainError("Sample not found",
		1000,
		fmt.Sprintf("No sample found for id %d", id))
}
