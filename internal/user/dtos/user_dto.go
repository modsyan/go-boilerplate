package dtos

import (
	"company-name/entities"
)

type UserDto struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

func UserDtoFromEntity(entity *entities.User) *UserDto {
	return &UserDto{
		ID:          entity.ID.Hex(),
		Email:       entity.Email,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		PhoneNumber: entity.PhoneNumber,
	}
}
