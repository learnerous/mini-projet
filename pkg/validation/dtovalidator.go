package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var instance *validator.Validate = nil

/*
this function is aimed to return a singleton instance of  *validator.Validate
the instance is preconfigured to return, when validation fails,the json name (the specified json name of the field in the struct tag) of the field on which we encountered a validation error
*/
func DtoValidator() *validator.Validate {
	if instance != nil {
		return instance
	}

	instance = validator.New()
	instance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//the default behavior of the validator is to return the name of the struct field (usually Capitalized) which is not what we want
		//the instruction below split the json Tag on ',', the json name will be at position 0
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return instance
}
