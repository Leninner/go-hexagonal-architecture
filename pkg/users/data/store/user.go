package users

import (
	"errors"

	"github.com/leninner/go-feature-flags/database"
	models "github.com/leninner/go-feature-flags/pkg/users/data/models"
	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
	"gorm.io/gorm"
)

const (
	userNotFound = "user not found"
)

// Store struct manages interactions with user store
type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&models.User{})

	return &Store{
		db: db,
	}
}

func (r *Store) GetByEmail(email string) (*entities.User, error) {
	user := &models.User{}

	database.DB.Where("email = ?", email).First(user)

	if user.ID == 0 {
		return nil, errors.New(userNotFound)
	}

	return user.ToDomainModel(user), nil
}

func (r *Store) GetUserModules(id int) (*entities.User, error) {
	user := &models.User{}

	database.DB.Where("id = ?", id).Preload("Role").First(user)

	if user.ID == 0 {
		return nil, errors.New(userNotFound)
	}

	return user.ToDomainModel(user), nil
}

func (r *Store) GetUserDetails(id int) (*entities.User, error) {
	user := &models.User{}
	database.DB.Where("id = ?", id).Preload("Role").First(user)

	if user.ID == 0 {
		return nil, errors.New(userNotFound)
	}

	return user.ToDomainModel(user), nil
}

func (r *Store) GetByID(id int) (*entities.User, error) {
	user := &models.User{}

	database.DB.Where("id = ?", id).First(user)

	if user.ID == 0 {
		return nil, errors.New(userNotFound)
	}

	return user.ToDomainModel(user), nil
}
