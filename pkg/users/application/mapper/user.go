package users

import (
	app "github.com/leninner/go-feature-flags/pkg/users/application/dto"
	domain "github.com/leninner/go-feature-flags/pkg/users/domain/entities"
)

func toResponseModel(entity *domain.User) *app.UserResponse {
	return &app.UserResponse{
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
