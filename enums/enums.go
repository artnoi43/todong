package enums

import "strings"

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

	// Capitalize to make in obvious in the code
	POSTGRES_MAX_STRLEN int = 65535
)

func (s Status) IsValid() bool {
	switch s.ToUpper() {
	case InProgress, Completed:
		return true
	}
	return false
}

func (s Status) ToUpper() Status {
	return Status(strings.ToUpper(string(s)))
}

func (s StoreType) IsValid() bool {
	switch s.ToUpper() {
	case Gorm, Redis:
		return true
	}
	return false
}

func (s StoreType) ToUpper() StoreType {
	return StoreType(strings.ToUpper(string(s)))
}

func (s ServerType) IsValid() bool {
	switch s.ToUpper() {
	case Gin, Fiber, Gorilla:
		return true
	}
	return false
}

func (s ServerType) ToUpper() ServerType {
	return ServerType(strings.ToUpper(string(s)))
}
