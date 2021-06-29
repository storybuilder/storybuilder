package error

import (
	"context"
	"net/http"

	baseErrs "github.com/storybuilder/storybuilder/app/errors"
	httpErrs "github.com/storybuilder/storybuilder/channels/http/errors"
	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	domainErrs "github.com/storybuilder/storybuilder/domain/errors"
	externalErrs "github.com/storybuilder/storybuilder/externals/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, logger adapters.LogAdapterInterface) (errMessage []byte, status int) {
	switch err.(type) {
	case *baseErrs.ServerError, *httpErrs.TransformerError:
		logger.Error(ctx, "Server Error", err)
		errMessage = format(err)
		status = http.StatusInternalServerError

	case *externalErrs.AdapterError, *httpErrs.MiddlewareError,
		*externalErrs.RepositoryError, *externalErrs.ServiceError,
		*domainErrs.DomainError:
		logger.Error(ctx, "Other Error", err)
		errMessage = format(err)
		status = http.StatusBadRequest

	case *httpErrs.ValidationError:
		logger.Error(ctx, "Unpacker Error", err)
		errMessage = format(err)
		status = http.StatusUnprocessableEntity

	default:
		logger.Error(ctx, "Unknown Error", err)
		errMessage = format(err)
		status = http.StatusInternalServerError
	}

	return errMessage, status
}

// HandleValidationErrors specifically handles validation errors thrown by the validator.
func HandleValidationErrors(ctx context.Context, errs map[string]string, logger adapters.LogAdapterInterface) (errMessage []byte, status int) {
	errMessage = formatValidationErrors(errs)

	logger.Error(ctx, "Validation Errors", string(errMessage))

	return errMessage, http.StatusUnprocessableEntity
}
