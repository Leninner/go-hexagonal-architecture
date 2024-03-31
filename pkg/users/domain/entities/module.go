package users

import (
	"time"
)

type Module struct {
	ID int

	Name           string
	ParentModuleID *int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
