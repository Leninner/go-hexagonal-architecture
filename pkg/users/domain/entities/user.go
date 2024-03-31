package users

import "time"

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Role     interface{}

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
