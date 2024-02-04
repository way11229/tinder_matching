package domain

import "errors"

var (
	// internal error
	ErrorInternalServerError = errors.New("internal server error")
	ErrorRecordNotFound      = errors.New("record not found")

	// invalidate parameter
	ErrorUserNameInvalidate                = errors.New("user name invalidate")
	ErrorUserHeightInvalidate              = errors.New("user height invalidate")
	ErrorUserGenderInvalidate              = errors.New("user gender invalidate")
	ErrorUserNumberOfWantedDatesInvalidate = errors.New("user number of wanted dates invalidate")
)
