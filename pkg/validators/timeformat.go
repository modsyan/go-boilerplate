package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func RegisterTimeFormatValidators(validate *validator.Validate) {
	validate.RegisterValidation("timeformat", func(fl validator.FieldLevel) bool {
		timeRegex := `^\d{2}:\d{2}:\d{2}$`
		matched, _ := regexp.MatchString(timeRegex, fl.Field().String())
		return matched
	})
}
