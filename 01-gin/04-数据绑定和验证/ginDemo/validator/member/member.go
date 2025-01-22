package member

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

func NameValid(v *validator.Validate, topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {
	if s, ok := field.Interface().(string); ok {
		if s == "admin" {
			return false
		}
	}
	return true
}
