package usecases

import "errors"

var (
	ErrUnableToSaveRecord   = errors.New("unable to save record")
	ErrRecordNotFound       = errors.New("unable to find record")
	ErrUnableToFetchRecords = errors.New("unable to fetch records")
)
