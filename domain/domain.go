package domain

const (
	SEARCH_LIMIT_DEFAULT = 10

	USER_NAME_LEN_MAX               = 100
	USER_HEIGHT_MAX                 = 250
	USER_NUMBER_OF_WANTED_DATES_MAX = 100
)

type ServiceManager struct {
	UserService UserService
}
