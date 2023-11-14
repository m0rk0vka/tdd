package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	v := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = v.Len()
		getField = v.Index
	case reflect.Struct:
		numberOfValues = v.NumField()
		getField = v.Field
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
