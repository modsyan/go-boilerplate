package dtos

import (
	"company-name/entities"
	"company-name/pkg/idgenerator"
	"company-name/pkg/validators"
)

type UpdateUserRequest struct {
	ID          string `json:"id" validate:"required"`
	FirstName   string `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name" validate:"required,min=2,max=50"`
	Password    string `json:"password" validate:"omitempty,min=8,max=100"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,min=11,max=11"`
}

func (req *UpdateUserRequest) ToEntity() (*entities.User, string, error) {
	userId, err := idgenerator.ToPersistenceID(req.ID)
	if err != nil {
		return nil, "", err
	}

	return &entities.User{
		ID:          userId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}, req.Password, nil
}

func (u *UpdateUserRequest) Validate(validator validators.IValidator) error {
	return validator.ValidateStruct(u)
}

type UpdateUserResponse struct {
	UserDto
}

func UpdateUserResponseFromEntity(user *entities.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		UserDto: *UserDtoFromEntity(user),
	}
}
