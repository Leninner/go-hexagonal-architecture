package users

import (
	"time"

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
