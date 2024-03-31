package users

import (
	"time"

	"gorm.io/gorm"
)

type RolOperations struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	RoleID      *int           `json:"roleID"`
	OperationID int            `json:"operationId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
