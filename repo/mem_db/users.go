package memdb

import "github.com/google/uuid"

type User struct {
	Id                  uuid.UUID
	Name                string
	Height              uint32
	RemainNumberOfDates uint32
}
