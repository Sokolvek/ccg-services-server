package usecases

import (
	"reflect"
)

func StructToInterface(obj interface{}) map[string]interface{} {

	v := reflect.ValueOf(obj)

	res := map[string]interface{}{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		res[field.Name] = v.Field(i).Interface()
	}

	return res
}
