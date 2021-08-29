package utils

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

func GetFieldsWithValue(doc interface{}) []string {
	var result []string

	s := reflect.TypeOf(doc).Elem()

	for i := 0; i < s.NumField(); i++ {
		n := strcase.ToSnake(s.Field(i).Name)
		e := reflect.ValueOf(doc).Elem().Field(i)
		switch e.Kind() {
		case reflect.Ptr:
			if !e.IsNil() {
				result = append(result, n)
			}
		case reflect.String:
			if len(s.String()) > 0 {
				result = append(result, n)
			}
		}
	}
	return result
}
