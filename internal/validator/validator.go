package validator

import (
	"strings"
	"unicode"
	"github.com/go-playground/validator/v10"
)
type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()

	v.RegisterValidation(
		"notblank",
		notBlank,
	)
	v.RegisterValidation(
		"strongpassword",
		strongPassword,
	)

	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func notBlank(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	return strings.TrimSpace(value) != ""
}

func strongPassword(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range value {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.IsSymbol(char) {
			hasSpecial = true
		}
		if unicode.IsPunct(char) {
			hasSpecial = true
		}
		if hasUpper && hasLower && hasDigit && hasSpecial {
			return true
		}
		if unicode.IsSpace(char) {
			return false
		}
	}
	return false
}