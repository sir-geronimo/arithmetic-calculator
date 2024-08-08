package usecases

import "errors"

var (
	ErrUnableToSaveOperation = errors.New("unable to save operation")
	ErrInvalidOperation      = errors.New("invalid operation type")
)
