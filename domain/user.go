package domain

import "github.com/google/uuid"

type UserGender string

const (
	USER_GENDER_MALE   UserGender = "male"
	USER_GENDER_FEMALE UserGender = "female"
)

type User struct {
	Id            uuid.UUID
	Name          string
	Height        uint32
	Gender        UserGender
	NumberOfDates uint32
	DateStatus    bool
}
