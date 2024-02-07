package validation

import "github.com/go-playground/validator/v10"

func ParseErrors(err error) []string {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	incorrectFields := make([]string, 0, len(validatorErrs))

	for _, e := range validatorErrs {
		incorrectFields = append(incorrectFields, e.Field())
	}

	return incorrectFields
}
