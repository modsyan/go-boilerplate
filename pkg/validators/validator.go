package validators

import (
	"company-name/pkg/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	ValidateStruct(s interface{}) *errors.BaseError
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator(validate *validator.Validate) IValidator {
	return &Validator{
		validate: validate,
	}
}

func (v *Validator) ValidateStruct(s interface{}) *errors.BaseError {
	err := v.validate.Struct(s)
	if err != nil {
		var validationErrors = make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			validationErrors[fieldErr.Field()] = fmt.Sprintf("validation failed on tag '%s'", fieldErr.ActualTag())
		}
		return errors.ValidationErrors(validationErrors)
	}
	return nil
}
