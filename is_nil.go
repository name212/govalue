// Copyright 2025
// license that can be found in the LICENSE file.

package govalue

import (
	"reflect"
)

func IsNil(value any) bool {
	iv := reflect.ValueOf(value)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}
