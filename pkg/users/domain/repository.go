package users

import (
	domain "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

type UserRepository interface {
	// GetByEmail returns a user by email
	GetByEmail(email string) (*domain.User, error)

	// GetByID returns a user by id
	GetByID(id int) (*domain.User, error)

	// GetUserModules returns a user by id
	GetUserModules(id int) (*domain.User, error)

	// GetUserDetails returns a user by id
	GetUserDetails(id int) (*domain.User, error)
}
