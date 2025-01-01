package errors

import (
	cons "company-name/constants/msgkey"
	"fmt"
	"net/http"

	"company-name/pkg/localization"
)

// HttpError defines an interface for HTTP-related errors providing status code, message, and error string methods.
type HttpError interface {
	StatusCode() int
	Message() string
	Error() string
}

// BaseError is a struct that contains an err message and an HTTP status code.
type BaseError struct {
	code             int
	message          string
	err              error
	validationErrors map[string]string
}

// StatusCode returns the HTTP status code associated with the BaseError instance.
func (e *BaseError) StatusCode() int {
	return e.code
}

// Message returns the error message contained in the BaseError instance.
func (e *BaseError) Message() string {
	return e.message
}

func (e *BaseError) ValidationErrors() map[string]string {
	return e.validationErrors
}

// Error returns the error message string of the BaseError. If an underlying error exists, its message is returned.
func (e *BaseError) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.message
}

// NewHTTPError creates a new BaseError.
func NewHTTPError(code int, message string) *BaseError {
	return &BaseError{
		code:    code,
		message: message,
	}
}

// NewHTTPErrorf creates a new BaseError with a formatted message.
func NewHTTPErrorf(code int, format string, args ...interface{}) *BaseError {
	return &BaseError{
		code:    code,
		message: fmt.Sprintf(format, args...),
	}
}

// NotFoundM creates a new BaseError with a 404 HTTP status code and a localized message using the given message key and error.
func NotFoundM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusNotFound, messageKey, err)
}

// NotFound creates a new BaseError with a 404 HTTP status code and a "not_found" message, wrapping the provided error.
func NotFound(err error) *BaseError {
	return NotFoundM(cons.ErrNotFound, err)
}

// UnauthorizedM creates a BaseError with a 401 status, a localized message based on messageKey, and an optional error.
func UnauthorizedM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusUnauthorized, messageKey, err)
}

// Unauthorized creates a new BaseError with a 401 unauthorized status code and an optional error for additional context.
func Unauthorized(err error) *BaseError {
	return UnauthorizedM(cons.ErrUnauthorized, err)
}

// ForbiddenM creates a new BaseError with an HTTP 403 Forbidden status, a localized message, and an optional error.
func ForbiddenM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusForbidden, messageKey, err)
}

// Forbidden creates a new BaseError with a 403 Forbidden status code and a localized error message.
func Forbidden(err error) *BaseError {
	return ForbiddenM(cons.ErrForbidden, err)
}

// BadRequestM creates a BaseError with a 400 status code, a localized message derived from messageKey, and the provided error.
func BadRequestM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusBadRequest, messageKey, err)
}

// BadRequest wraps an error into a BaseError with a 400 status code and a localized "bad_request" message.
func BadRequest(err error) *BaseError {
	return BadRequestM(cons.ErrBadRequest, err)
}

// InternalServerErrorM creates a BaseError with a 500 status code, a localized message, and the provided error details.
func InternalServerErrorM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusInternalServerError, messageKey, err)
}

// InternalServerError generates a BaseError with a 500 status code and an associated error for internal server issues.
func InternalServerError(err error) *BaseError {
	return InternalServerErrorM(cons.ErrInternalServerError, err)
}

// ConflictM creates a new BaseError with a 409 HTTP status code using the provided message key and error details.
func ConflictM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusConflict, messageKey, err)
}

// Conflict creates a new BaseError with a 409 HTTP status code and the provided error details.
func Conflict(err error) *BaseError {
	return ConflictM(cons.ErrConflict, err)
}

// UnprocessableEntityM creates a BaseError with a 422 HTTP status code and a localized error message using a message key.
func UnprocessableEntityM(messageKey string, err error) *BaseError {
	return NewLocalizedHTTPError(http.StatusUnprocessableEntity, messageKey, err)
}

// UnprocessableEntity creates a BaseError with a 422 HTTP status code and a localized unprocessable entity error message.
func UnprocessableEntity(err error) *BaseError {
	return UnprocessableEntityM(cons.ErrUnprocessableEntity, err)
}

// ValidationErrors generates a BaseError with validation errors mapped to a 422 status code and localized message.
func ValidationErrors(errors map[string]string) *BaseError {
	return ValidationErrorsM(cons.ErrValidationFailed, errors)
}

// ValidationErrorsM creates a new BaseError with a 422 status code, localized message, and map of validation errors.
func ValidationErrorsM(message string, errors map[string]string) *BaseError {
	return &BaseError{
		code:             http.StatusUnprocessableEntity,
		message:          localization.L(message),
		validationErrors: errors,
	}
}
