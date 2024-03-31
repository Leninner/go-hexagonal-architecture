package users

import (
	"time"

	entities "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password" gorm:"->:false"`
	Email    string `json:"email" gorm:"unique"`

	Role   Role `json:"role" gorm:"foreignKey:RoleID"`
	RoleID int  `json:"roleId"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (u *User) ToDBModel(entity *entities.User) *User {
	return &User{
		Username: entity.Username,
	}
}

func (u *User) ToDomainModel(entity *User) *entities.User {
	return &entities.User{
		ID:        entity.ID,
		Username:  entity.Username,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
