package dtos

import (
	"company-name/entities"
	"company-name/pkg/idgenerator"
	"company-name/pkg/validators"
	"time"
)

type CreateUserRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=128"`
	FirstName   string `json:"first_name" validate:"required,min=2,max=50"`
	LastName    string `json:"last_name" validate:"required,min=2,max=50"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Role        string `json:"role" validate:"required,oneof=admin user moderator"`
}

func (req *CreateUserRequest) Validate(validate validators.IValidator) error {
	return validate.ValidateStruct(req)
}

func (req *CreateUserRequest) ToEntity() (user *entities.User, password string) {
	user = &entities.User{
		ID:          idgenerator.GenerateID(),
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	password = req.Password
	return
}

type CreateUserResponse struct {
	UserDto
}

func CreateUserResponseFromEntity(user *entities.User) *CreateUserResponse {
	return &CreateUserResponse{
		UserDto: *UserDtoFromEntity(user),
	}
}
