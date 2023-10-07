package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		var validationErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			var errorMessage string
			fieldName := e.Field()
			
			field, found := reflect.TypeOf(i).FieldByName(fieldName)
			if found {
				jsonTag := field.Tag.Get("json")
				if jsonTag != "" && jsonTag != "-" {
					fieldName = jsonTag
				}
			}
			
			switch e.Tag() {
			case "required":
				errorMessage = fmt.Sprintf("%s is required", fieldName)
			case "max":
				errorMessage = fmt.Sprintf("%s should be less than or equal to %s", fieldName, e.Param())
			case "min":
				errorMessage = fmt.Sprintf("%s should be greater than or equal to %s", fieldName, e.Param())
			case "email":
				errorMessage = fmt.Sprintf("%s should be a valid email", fieldName)
			case "oneof":
				value := reflect.ValueOf(i).FieldByName(e.Field()).Interface()
				allowedValues := strings.Split(e.Param(), " ")
				
				valid := false
				for _, v := range allowedValues {
					if value == v {
						valid = true
						break
					}
				}
				
				if !valid {
					errorMessage = fmt.Sprintf("%s is not one of the allowed values: %s", fieldName, allowedValues)
				}
			default:
				errorMessage = fmt.Sprintf("%s is %s", fieldName, e.Tag())
			}
			validationErrors = append(validationErrors, errorMessage)
		}
		
		errorMessage := strings.Join(validationErrors, ", ")
		
		return fmt.Errorf(errorMessage)
	}
	
	return nil
}
