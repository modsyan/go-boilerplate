package dtos

import (
	"company-name/constants"
	"company-name/entities"
	"company-name/pkg/idgenerator"
	"time"
)

type RegisterRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func (d *RegisterRequest) ToEntity(hashedPassword string) entities.User {
	return entities.User{
		ID:             idgenerator.GenerateID(),
		FirstName:      d.FirstName,
		LastName:       d.LastName,
		Email:          d.Email,
		HashedPassword: hashedPassword,
		PhoneNumber:    d.PhoneNumber,
		Status:         constants.UserStatusPending, // Default status during registration
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

type RegisterResponse struct {
	UserStatus      string `json:"user_status" example:"pending"`
	VerificationURL string `json:"verification_url,omitempty"`
}

func (d *RegisterResponse) FromEntity(user *entities.User, verificationURL string) {
	d.UserStatus = user.Status
	d.VerificationURL = verificationURL
}
