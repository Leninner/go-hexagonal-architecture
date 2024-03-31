package users

import (
	"time"

	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
	"gorm.io/gorm"
)

type Module struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name"`
	ParentModuleID *int           `json:"parentModuleId"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (m *Module) ToDBModel(entity *entities.Module) *Module {
	return &Module{
		ID:             entity.ID,
		Name:           entity.Name,
		ParentModuleID: entity.ParentModuleID,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}

func (m *Module) ToDomainModel(entity *Module) *entities.Module {
	return &entities.Module{
		ID:             entity.ID,
		Name:           entity.Name,
		ParentModuleID: entity.ParentModuleID,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}
}
