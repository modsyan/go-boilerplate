package errors

import "company-name/pkg/localization"

type LocalizedError struct {
	*BaseError
	messageKey string
}

func NewLocalizedError(code int, messageKey string, err error) *LocalizedError {
	return &LocalizedError{
		BaseError:  NewHTTPError(code, localization.L(messageKey)),
		messageKey: messageKey,
	}
}

func (e *LocalizedError) Localize() string {
	return localization.L(e.messageKey)
}

func NewLocalizedHTTPError(code int, messageKey string, err error) *BaseError {
	return &BaseError{
		code:    code,
		message: localization.L(messageKey),
		err:     err,
	}
}
