package entities

import (
	"company-name/pkg/validators"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockKey struct {
	Page    string `json:"page" validate:"required" index:"page_section_idx,unique" bson:"page"`
	Section string `json:"section" validate:"required" index:"page_section_idx,unique" bson:"section"`
}

type ContentBlocks struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Key       BlockKey           `bson:"key" json:"key"`                             // Unique key for the content block
	Content   string             `bson:"content" json:"content" validate:"required"` // Content of the block Can be in HTML or Markdown or Plain Text
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func (s *ContentBlocks) Validate(validator validators.IValidator) error {
	if err := validator.ValidateStruct(s); err != nil {
		return err
	}
	return nil
}
