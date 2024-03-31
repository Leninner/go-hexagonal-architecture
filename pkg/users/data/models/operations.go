package users

import (
	"time"

	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
	"gorm.io/gorm"
)

type Operation struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ModuleID    int    `json:"module_id"`
	Module      Module `json:"module" gorm:"foreignKey:ModuleID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *Operation) ToDBModel(entity *entities.Operation) *Operation {
	return &Operation{
		Name:        entity.Name,
		Description: entity.Description,
		ModuleID:    entity.ModuleID,
	}
}

func (o *Operation) ToDomainModel(entity *Operation) *entities.Operation {
	return &entities.Operation{
		Name:        entity.Name,
		Description: entity.Description,
		ModuleID:    entity.ModuleID,
	}
}
