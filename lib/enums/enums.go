package enums

import "github.com/artnoi43/mgl/str"

type (
	Status     string
	StoreType  string
	ServerType string
)

// Go standard library uses PascalCase/camelCase to name constants,
// except for something like os.O_RDONLY which is directly referencing POSIX
const (
	// Only 2 status messages are valid
	InProgress Status = "IN_PROGRESS"
	Completed  Status = "COMPLETED"

	// Data storage enum
	Gorm  StoreType = "GORM"
	Redis StoreType = "REDIS"

	// HTTP web framework
	Gin     ServerType = "GIN"
	Fiber   ServerType = "FIBER"
	Gorilla ServerType = "GORILLA"
	NetHttp ServerType = "NETHTTP"

	// Capitalize to make in obvious in the code
	POSTGRES_MAX_STRLEN int = 65535
)

func (s Status) IsValid() bool {
	switch str.ToUpper(s) {
	case InProgress, Completed:
		return true
	}
	return false
}

func (s StoreType) IsValid() bool {
	switch str.ToUpper(s) {
	case Gorm, Redis:
		return true
	}
	return false
}

func (s ServerType) IsValid() bool {
	switch str.ToUpper(s) {
	case Gin, Fiber, Gorilla, NetHttp:
		return true
	}
	return false
}
