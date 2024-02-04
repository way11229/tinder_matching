package domain

import "errors"

var (
	// internal error
	ErrorInternalServerError = errors.New("internal server error")
	ErrorRecordNotFound      = errors.New("record not found")

	// invalidate parameter
	ErrorMissRequiredParameters         = errors.New("miss required parameters")
	ErrorUserIdInvalid                  = errors.New("user id is invalid")
	ErrorUserNameInvalid                = errors.New("user name is invalid")
	ErrorUserHeightInvalid              = errors.New("user height is invalid")
	ErrorUserGenderInvalid              = errors.New("user gender si invalid")
	ErrorUserNumberOfWantedDatesInvalid = errors.New("user number of wanted dates is invalid")
)
