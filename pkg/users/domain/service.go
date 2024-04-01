package users

import (
	"strconv"

	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

type UserService interface {
	// GetUser returns a user by its ID
	GetUser(id string) (*entities.User, error)

	// GetByEmail returns a user by its email
	GetByEmail(email string) (*entities.User, error)
}

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUser(id string) (*entities.User, error) {
	integerID, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	return s.repository.GetByID(integerID)
}

func (s *Service) GetByEmail(email string) (*entities.User, error) {
	return s.repository.GetByEmail(email)
}
