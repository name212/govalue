// Copyright 2026
// license that can be found in the LICENSE file.

package govalue

import (
	"errors"
	"fmt"
	"reflect"
)

var CannotConvertErr = errors.New("cannot convert")

func ConvertTo[F any, T any](v F) (T, error) {
	var zero T

	fType := reflect.TypeFor[F]()
	if fType.Kind() != reflect.Interface {
		return zero, fmt.Errorf("%w %s is not interface", CannotConvertErr, fType.String())
	}

	tType := reflect.TypeFor[T]()
	if tType.Kind() != reflect.Pointer {
		return zero, fmt.Errorf("%w %s is not pointer", CannotConvertErr, tType.String())
	}

	if !tType.Implements(fType) {
		return zero, fmt.Errorf("%w from %s to %s", CannotConvertErr, fType.String(), tType.String())
	}

	return any(v).(T), nil
}
