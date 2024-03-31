package users

import (
	"time"

	domain "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

// AuthorResponse struct defines response fields
type UserResponse struct {
	ID         int            `json:"id"`
	Username   string         `json:"username"`
	Email      string         `json:"email"`
	Role       interface{}    `json:"role"`
	Modules    *interface{}   `json:"modules"`
	CreatedAt  time.Time      `json:"createdAt"`
	Operations *[]interface{} `json:"operations"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  time.Time      `json:"deletedAt"`
}

// ListResponse struct defines authors list response structure
type ListResponse struct {
	Data []UserResponse `json:"data"`
}

func (u *UserResponse) ToResponseModel(entity *domain.User) *UserResponse {
	return &UserResponse{
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
