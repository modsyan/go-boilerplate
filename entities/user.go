package entities

import (
	"company-name/pkg/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Email          string             `bson:"email" json:"email" validate:"required"`
	HashedPassword string             `bson:"hashed_password" json:"hashed_password" validate:"required"`
	FirstName      string             `bson:"first_name" json:"first_name" validate:"required"`
	LastName       string             `bson:"last_name" json:"last_name" validate:"required"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number" validate:"required"`
	Status         string             `bson:"status" json:"status" validate:"required" example:"pending"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

func (s *User) Validate(validator validators.IValidator) error {
	if err := validator.ValidateStruct(s); err != nil {
		return err
	}
	return nil
}
