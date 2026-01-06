// Copyright 2025
// license that can be found in the LICENSE file.

package govalue

import (
	"reflect"
)

// Nil
// returns true if value is nil for any type
// and true if interface value is not nil but have nil data
func Nil(value any) bool {
	iv := reflect.ValueOf(value)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Interface, reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Chan:
		return iv.IsNil()
	default:
		return false
	}
}

// IsNil
// Deprecated:
// Use Nil instead of IsNil
func IsNil(value any) bool {
	return Nil(value)
}

// NotNil
// returns inversion of Nil function
func NotNil(value any) bool {
	return !Nil(value)
}
