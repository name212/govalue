// Copyright 2025
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"testing"

	"github.com/name212/govalue"
	"github.com/stretchr/testify/require"
)

type testInterface interface {
	Do()
}

type testInterfaceImpl struct{}

func (*testInterfaceImpl) Do() {
	fmt.Println("Call Do on testInterfaceImpl")
}

type testErrorImpl struct{}

func (*testErrorImpl) Error() string {
	return "error"
}

func TestIsNil(t *testing.T) {
	require.True(t, govalue.IsNil(nil))

	var ii *testInterfaceImpl
	require.True(t, govalue.IsNil(ii))

	var err error
	require.True(t, govalue.IsNil(err))

	var ival *int
	require.True(t, govalue.IsNil(ival))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.True(t, govalue.IsNil(actual))
	}

	var v testInterface
	require.True(t, govalue.IsNil(v))
	assertCallFuncWithInterface(t, v)

	var m map[string]struct{}
	require.True(t, govalue.IsNil(m))

	m = nil
	require.True(t, govalue.IsNil(m))

	var s []string
	require.True(t, govalue.IsNil(s))

	s = nil
	require.True(t, govalue.IsNil(s))

	var f func()
	require.True(t, govalue.IsNil(f))

	f = nil
	require.True(t, govalue.IsNil(f))

	var ch chan struct{}
	require.True(t, govalue.IsNil(ch))

	ch = nil
	require.True(t, govalue.IsNil(ch))

	returnsNil := func() testInterface {
		return nil
	}
	require.True(t, govalue.IsNil(returnsNil()))

	vv := returnsNil()
	require.True(t, govalue.IsNil(vv))

	returnsNilFromLocal := func() testInterface {
		var i testInterface
		return i
	}
	require.True(t, govalue.IsNil(returnsNilFromLocal()))

	vvv := returnsNilFromLocal()
	require.True(t, govalue.IsNil(vvv))

	returnsNilFromLocalPointer := func() testInterface {
		var i *testInterfaceImpl
		return i
	}
	require.True(t, govalue.IsNil(returnsNilFromLocalPointer()))

	vvvv := returnsNilFromLocalPointer()
	require.True(t, govalue.IsNil(vvvv))

	returnsNilPointer := func() *testInterfaceImpl {
		var i *testInterfaceImpl
		return i
	}
	require.True(t, govalue.IsNil(returnsNilPointer()))

	vvvvv := returnsNilPointer()
	require.True(t, govalue.IsNil(vvvvv))

	returnsErr := func() error {
		return nil
	}
	require.True(t, govalue.IsNil(returnsErr()))

	err1 := returnsErr()
	require.True(t, govalue.IsNil(err1))

	returnsErrImpl := func() error {
		var e *testErrorImpl
		return e
	}
	require.True(t, govalue.IsNil(returnsErrImpl()))

	err2 := returnsErrImpl()
	require.True(t, govalue.IsNil(err2))

}

func TestIsNotNil(t *testing.T) {
	require.False(t, govalue.IsNil(struct{}{}))
	require.False(t, govalue.IsNil(""))
	require.False(t, govalue.IsNil(0))
	require.False(t, govalue.IsNil(0.0))
	require.False(t, govalue.IsNil('a'))
	require.False(t, govalue.IsNil(testInterfaceImpl{}))
	require.False(t, govalue.IsNil(&testInterfaceImpl{}))

	var err error
	err = fmt.Errorf("test error")
	require.False(t, govalue.IsNil(err))

	i := testInterfaceImpl{}
	require.False(t, govalue.IsNil(i))

	ii := &testInterfaceImpl{}
	require.False(t, govalue.IsNil(ii))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.False(t, govalue.IsNil(actual))
	}

	var v testInterface
	v = &testInterfaceImpl{}
	require.False(t, govalue.IsNil(v))
	assertCallFuncWithInterface(t, v)

	m := make(map[string]struct{})
	require.False(t, govalue.IsNil(m))

	s := make([]string, 0, 0)
	require.False(t, govalue.IsNil(s))

	f := func() {}
	require.False(t, govalue.IsNil(f))

	ch := make(chan struct{})
	require.False(t, govalue.IsNil(ch))

	returnsNoneNilFromLocal := func() testInterface {
		var iii testInterface
		iii = &testInterfaceImpl{}
		return iii
	}
	require.False(t, govalue.IsNil(returnsNoneNilFromLocal()))

	vv := returnsNoneNilFromLocal()
	require.False(t, govalue.IsNil(vv))

	returnsNoneNil := func() testInterface {
		return &testInterfaceImpl{}
	}
	require.False(t, govalue.IsNil(returnsNoneNil()))

	vvv := returnsNoneNil()
	require.False(t, govalue.IsNil(vvv))

	returnsNoneNilErr := func() error {
		return fmt.Errorf("test error")
	}
	require.False(t, govalue.IsNil(returnsNoneNilErr()))

	err1 := returnsNoneNilErr()
	require.False(t, govalue.IsNil(err1))

	returnsNoneNilErrImpl := func() error {
		return &testErrorImpl{}
	}
	require.False(t, govalue.IsNil(returnsNoneNilErrImpl()))

	err2 := returnsNoneNilErrImpl()
	require.False(t, govalue.IsNil(err2))

	returnsNoneNilStruct := func() *testInterfaceImpl {
		return &testInterfaceImpl{}
	}
	require.False(t, govalue.IsNil(returnsNoneNilStruct()))

	vs := returnsNoneNilStruct()
	require.False(t, govalue.IsNil(vs))

	returnsStruct := func() testInterfaceImpl {
		return testInterfaceImpl{}
	}
	require.False(t, govalue.IsNil(returnsStruct()))

	ns := returnsStruct()
	require.False(t, govalue.IsNil(ns))
}

func TestIsNilAfterSet(t *testing.T) {
	err := fmt.Errorf("test error")
	err = nil
	require.True(t, govalue.IsNil(err))

	ii := &testInterfaceImpl{}
	ii = nil
	require.True(t, govalue.IsNil(ii))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.True(t, govalue.IsNil(actual))
	}

	var v testInterface = &testInterfaceImpl{}
	v = nil
	require.True(t, govalue.IsNil(v))
	assertCallFuncWithInterface(t, v)

	m := make(map[string]struct{})
	m = nil
	require.True(t, govalue.IsNil(m))

	s := make([]string, 0, 0)
	s = nil
	require.True(t, govalue.IsNil(s))

	f := func() {}
	f = nil
	require.True(t, govalue.IsNil(f))

	ch := make(chan struct{})
	ch = nil
	require.True(t, govalue.IsNil(ch))
}
