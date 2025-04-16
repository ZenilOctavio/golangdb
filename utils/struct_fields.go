package utils

import (
	"db/constants"
	"errors"
	"reflect"
)

func isStruct(s interface{}) bool {
	if s == nil {
		return false
	}

	t := reflect.TypeOf(s)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	return t.Kind() == reflect.Struct
}

type GetStructFieldsError int

const (
	NOT_AN_STRUCT GetStructFieldsError = iota
)

var errorsMapper = []string{
	"The given element %v is not a struct",
}

func GetStructFields[T any](v T) (map[string]string, error) {

	if !isStruct(v) {
		return nil, errors.New(errorsMapper[NOT_AN_STRUCT])
	}

	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	attsToFieldsMapper := make(map[string]string)

	for i := range t.NumField() {
		field := t.Field(i)

		if field.PkgPath != "" {
			continue
		}

		fieldName := field.Tag.Get(constants.FIELD_NAME_TAG)

		if fieldName == "" {
			continue
		}

		name := field.Name

		attsToFieldsMapper[name] = fieldName
	}

	return attsToFieldsMapper, nil

}
