// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"testing"

	"github.com/name212/govalue"
	"github.com/stretchr/testify/require"
)

type i interface {
	Do()
}

type impl struct{}

func (impl) Do() {}

type another interface {
	DoAnother()
}

type anotherImpl struct{}

func (anotherImpl) DoAnother() {}

func TestConvertTo(t *testing.T) {
	var ii i = &impl{}
	res, err := govalue.ConvertTo[i, *impl](ii)

	require.NoError(t, err)
	require.IsType(t, res, &impl{})
}

func TestConvertToFail(t *testing.T) {
	a := &anotherImpl{}
	res, err := govalue.ConvertTo[another, *impl](a)

	require.Error(t, err)
	require.ErrorIs(t, err, govalue.CannotConvertErr)
	require.Nil(t, res)
}

func TestConvertFromNotInterface(t *testing.T) {
	a := &anotherImpl{}
	res, err := govalue.ConvertTo[another, *impl](a)

	require.Error(t, err)
	require.ErrorIs(t, err, govalue.CannotConvertErr)
	require.Nil(t, res)
}

func TestConvertToNotPointer(t *testing.T) {
	a := impl{}
	_, err := govalue.ConvertTo[i, impl](a)

	require.Error(t, err)
	require.ErrorIs(t, err, govalue.CannotConvertErr)
}
