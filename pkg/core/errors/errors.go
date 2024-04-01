package core

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leninner/go-feature-flags/pkg/core"
)

const (
	// NotFound error indicates a missing / not found record
	NotFound        = "NotFound"
	notFoundMessage = "Record not found"

	// ValidationError indicates an error in input validation
	ValidationError        = "ValidationError"
	validationErrorMessage = "Validation error"

	// ResourceAlreadyExists indicates a duplicate / already existing record
	ResourceAlreadyExists     = "ResourceAlreadyExists"
	alreadyExistsErrorMessage = "Resource already exists"

	// RepositoryError indicates a repository (e.g database) error
	RepositoryError        = "RepositoryError"
	repositoryErrorMessage = "Error in repository operation"

	// NotAuthenticated indicates an authentication error
	NotAuthenticated             = "NotAuthenticated"
	notAuthenticatedErrorMessage = "Not Authenticated"

	// TokenGeneratorError indicates an token generation error
	TokenGeneratorError        = "TokenGeneratorError"
	tokenGeneratorErrorMessage = "Error in token generation"

	// NotAuthorized indicates an authorization error
	NotAuthorized             = "NotAuthorized"
	notAuthorizedErrorMessage = "Not Authorized"

	// UnknownError indicates an error that the app cannot find the cause for
	UnknownError        = "UnknownError"
	unknownErrorMessage = "Something went wrong"
)

// AppError defines an application (domain) error
type AppError struct {
	Err  error
	Type string
}

// NewAppError initializes a new domain error using an error and its type.
func NewAppError(err error, errType string) *AppError {
	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// NewAppErrorWithType initializes a new default error for a given type.
func NewAppErrorWithType(errType string) *AppError {
	var err error

	switch errType {
	case NotFound:
		err = errors.New(notFoundMessage)
	case ValidationError:
		err = errors.New(validationErrorMessage)
	case ResourceAlreadyExists:
		err = errors.New(alreadyExistsErrorMessage)
	case RepositoryError:
		err = errors.New(repositoryErrorMessage)
	case NotAuthenticated:
		err = errors.New(notAuthenticatedErrorMessage)
	case NotAuthorized:
		err = errors.New(notAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = errors.New(tokenGeneratorErrorMessage)
	default:
		err = errors.New(unknownErrorMessage)
	}

	return &AppError{
		Err:  err,
		Type: errType,
	}
}

// String converts the app error to a human-readable string.
func (appErr *AppError) Error() string {
	return appErr.Err.Error()
}

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs.Last().Err.(*AppError)

		if ok {
			errorDescription := core.NewResponseMessage(err.Error(), nil)

			switch err.Type {
			case NotFound:
				c.JSON(http.StatusNotFound, errorDescription)
				return
			case ValidationError:
				c.JSON(http.StatusBadRequest, errorDescription)
				return
			case ResourceAlreadyExists:
				c.JSON(http.StatusConflict, errorDescription)
				return
			case NotAuthenticated:
				c.JSON(http.StatusUnauthorized, errorDescription)
				return
			case NotAuthorized:
				c.JSON(http.StatusForbidden, errorDescription)
				return
			case RepositoryError:
			default:
				c.JSON(http.StatusInternalServerError, errorDescription)
				return
			}
		}

		c.JSON(http.StatusInternalServerError, core.NewResponseMessage(errs.Last().Error(), nil))
		return
	}
}
