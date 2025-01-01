package dtos

import (
	"company-name/entities"
)

type GetUserDetailsRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetUserDetailsResponse struct {
	UserDto
}

func GetUserDetailsResponseFromEntity(user *entities.User) *GetUserDetailsResponse {
	return &GetUserDetailsResponse{
		UserDto: *UserDtoFromEntity(user),
	}
}
