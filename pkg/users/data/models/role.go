package users

import (
	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (r *Role) ToDBModel(entity *entities.Role) *Role {
	return &Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func (r *Role) ToDomainModel(entity *Role) *entities.Role {
	return &entities.Role{
		Name: entity.Name,
	}
}
