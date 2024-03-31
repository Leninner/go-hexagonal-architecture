package users

import (
	"strconv"

	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

// UserService defines author service behavior.
type UserService interface {
	GetUser(id string) (*entities.User, error)

	GetByEmail(email string) (*entities.User, error)
}

// Service struct handles author business logic tasks.
type Service struct {
	repository UserRepository
}

// GetUser returns an author by its id.
func (s *Service) GetUser(id string) (*entities.User, error) {
	integerID, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return s.repository.GetByID(integerID)
}

// GetByEmail returns an author by its email.
func (s *Service) GetByEmail(email string) (*entities.User, error) {
	return s.repository.GetByEmail(email)

}

// NewService creates a new service struct
func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}
