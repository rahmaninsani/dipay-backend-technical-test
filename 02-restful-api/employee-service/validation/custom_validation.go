package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage string
			switch e.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", e.Field())
			case "max":
				errorMessage = fmt.Sprintf("%s should be less than or equal to %s", e.Field(), e.Param())
			case "min":
				errorMessage = fmt.Sprintf("%s should be greater than or equal to %s", e.Field(), e.Param())
			default:
				errorMessage = fmt.Sprintf("%s is %s", e.Field(), e.Tag())
			}
			validationErrors = append(validationErrors, errorMessage)
		}
		
		errorMessage := strings.Join(validationErrors, ", ")
		
		return fmt.Errorf(errorMessage)
	}
	
	return nil
}
