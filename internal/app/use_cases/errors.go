package usecases

import "errors"

var (
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrUnableToSaveRecord        = errors.New("unable to save record")
	ErrRecordNotFound            = errors.New("unable to find record")
	ErrUnableToFetchRecords      = errors.New("unable to fetch records")
	ErrUnableToFindRecord        = errors.New("unable to find record")
	ErrUnableToPerformOperation  = errors.New("unable to perform operation")
	ErrOperationNotFound         = errors.New("unable to find operation")
	ErrUnableToSaveOperation     = errors.New("unable to save operation")
	ErrInvalidOperationType      = errors.New("invalid operation type")
	ErrInsufficientBalance       = errors.New("cannot perform operation. Insufficient balance")
	ErrOperationAlreadyPerformed = errors.New("cannot reprocess operation. Operation already performed")
)
