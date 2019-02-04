package slice

import (
	"reflect"
)

func Where(data interface{}, where func(i int) bool, sel func(i int)) {
	rv := reflect.ValueOf(data)
	length := rv.Len()
	for i := 0; i < length; i++ {
		if where(i) {
			sel(i)
		}
	}
}

func Contains(data interface{}, where func(i int) bool) bool {
	rv := reflect.ValueOf(data)
	length := rv.Len()
	for i := 0; i < length; i++ {
		if where(i) {
			return true
		}
	}
	return false
}

func Index(data interface{}, where func(i int) bool) int {
	rv := reflect.ValueOf(data)
	length := rv.Len()
	for i := 0; i < length; i++ {
		if where(i) {
			return i
		}
	}
	return -1
}
