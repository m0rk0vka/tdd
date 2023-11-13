package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	v := getValue(x)

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
