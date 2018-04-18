package cyclic

import (
	"reflect"
	"unsafe"
)

func IsCyclic(x interface{}) bool {
	seen := make([]unsafe.Pointer, 0)
	return isCyclic(reflect.ValueOf(x), seen)
}

func hasSeen(ptr unsafe.Pointer, seen []unsafe.Pointer) bool {
	for _, saw := range seen {
		if saw == ptr {
			return true
		}
	}
	return false
}

func isCyclic(x reflect.Value, seen []unsafe.Pointer) bool {

	if !x.IsValid() {
		return false
	}

	if x.CanAddr() && x.Kind() != reflect.Struct && x.Kind() != reflect.Array {
		xptr := unsafe.Pointer(x.UnsafeAddr())

		if hasSeen(xptr, seen) {
			return true
		}
		seen = append(seen, xptr)
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCyclic(x.Elem(), seen)

	case reflect.Slice, reflect.Array:
		for i := 0; i < x.Len(); i++ {
			if isCyclic(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCyclic(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCyclic(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}
	return false
}
