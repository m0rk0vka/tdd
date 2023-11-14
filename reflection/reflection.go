package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	v := getValue(x)

	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch v.Kind() {
	case reflect.String:
		fn(v.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walkValue(v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walkValue(v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walkValue(v.MapIndex(key))
		}
	case reflect.Chan:
		for x, ok := v.Recv(); ok; x, ok = v.Recv() {
			walkValue(x)
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
