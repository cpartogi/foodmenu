package utils

import (
	"fmt"

	"github.com/cpartogi/foodmenu/pkg/helper"
	"github.com/go-playground/validator/v10"
)

func errorType(err error) (int, error) {
	return helper.CommonError(err)
}

func switchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field := SetLowerAndAddSpace(err.Field())

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is mandatory",
					field)
			case "number":
				message = fmt.Sprintf("%s must be numbers only",
					field)
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					field, err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					field, err.Param())
			default:
				message = err.Error()
			}

			break
		}
	}
	return
}
