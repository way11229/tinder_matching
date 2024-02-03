package domain

import "errors"

var (
	ErrorInternalServerError = errors.New("internal server error")
	ErrorRecordNotFound      = errors.New("record not found")
)
