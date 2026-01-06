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

func TestNil(t *testing.T) {
	require.True(t, govalue.Nil(nil))

	var ii *testInterfaceImpl
	require.True(t, govalue.Nil(ii))

	var err error
	require.True(t, govalue.Nil(err))

	var ival *int
	require.True(t, govalue.Nil(ival))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.True(t, govalue.Nil(actual))
	}

	var v testInterface
	require.True(t, govalue.Nil(v))
	assertCallFuncWithInterface(t, v)

	var m map[string]struct{}
	require.True(t, govalue.Nil(m))

	m = nil
	require.True(t, govalue.Nil(m))

	var s []string
	require.True(t, govalue.Nil(s))

	s = nil
	require.True(t, govalue.Nil(s))

	var f func()
	require.True(t, govalue.Nil(f))

	f = nil
	require.True(t, govalue.Nil(f))

	var ch chan struct{}
	require.True(t, govalue.Nil(ch))

	ch = nil
	require.True(t, govalue.Nil(ch))

	returnsNil := func() testInterface {
		return nil
	}
	require.True(t, govalue.Nil(returnsNil()))

	vv := returnsNil()
	require.True(t, govalue.Nil(vv))

	returnsNilFromLocal := func() testInterface {
		var i testInterface
		return i
	}
	require.True(t, govalue.Nil(returnsNilFromLocal()))

	vvv := returnsNilFromLocal()
	require.True(t, govalue.Nil(vvv))

	returnsNilFromLocalPointer := func() testInterface {
		var i *testInterfaceImpl
		return i
	}
	require.True(t, govalue.Nil(returnsNilFromLocalPointer()))

	vvvv := returnsNilFromLocalPointer()
	require.True(t, govalue.Nil(vvvv))

	returnsNilPointer := func() *testInterfaceImpl {
		var i *testInterfaceImpl
		return i
	}
	require.True(t, govalue.Nil(returnsNilPointer()))

	vvvvv := returnsNilPointer()
	require.True(t, govalue.Nil(vvvvv))

	returnsErr := func() error {
		return nil
	}
	require.True(t, govalue.Nil(returnsErr()))

	err1 := returnsErr()
	require.True(t, govalue.Nil(err1))

	returnsErrImpl := func() error {
		var e *testErrorImpl
		return e
	}
	require.True(t, govalue.Nil(returnsErrImpl()))

	err2 := returnsErrImpl()
	require.True(t, govalue.Nil(err2))

}

func TestIsNotNil(t *testing.T) {
	require.False(t, govalue.Nil(struct{}{}))
	require.False(t, govalue.Nil(""))
	require.False(t, govalue.Nil(0))
	require.False(t, govalue.Nil(0.0))
	require.False(t, govalue.Nil('a'))
	require.False(t, govalue.Nil(testInterfaceImpl{}))
	require.False(t, govalue.Nil(&testInterfaceImpl{}))

	var err error
	err = fmt.Errorf("test error")
	require.False(t, govalue.Nil(err))

	i := testInterfaceImpl{}
	require.False(t, govalue.Nil(i))

	ii := &testInterfaceImpl{}
	require.False(t, govalue.Nil(ii))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.False(t, govalue.Nil(actual))
	}

	var v testInterface
	v = &testInterfaceImpl{}
	require.False(t, govalue.Nil(v))
	assertCallFuncWithInterface(t, v)

	m := make(map[string]struct{})
	require.False(t, govalue.Nil(m))

	s := make([]string, 0, 0)
	require.False(t, govalue.Nil(s))

	f := func() {}
	require.False(t, govalue.Nil(f))

	ch := make(chan struct{})
	require.False(t, govalue.Nil(ch))

	returnsNoneNilFromLocal := func() testInterface {
		var iii testInterface
		iii = &testInterfaceImpl{}
		return iii
	}
	require.False(t, govalue.Nil(returnsNoneNilFromLocal()))

	vv := returnsNoneNilFromLocal()
	require.False(t, govalue.Nil(vv))

	returnsNoneNil := func() testInterface {
		return &testInterfaceImpl{}
	}
	require.False(t, govalue.Nil(returnsNoneNil()))

	vvv := returnsNoneNil()
	require.False(t, govalue.Nil(vvv))

	returnsNoneNilErr := func() error {
		return fmt.Errorf("test error")
	}
	require.False(t, govalue.Nil(returnsNoneNilErr()))

	err1 := returnsNoneNilErr()
	require.False(t, govalue.Nil(err1))

	returnsNoneNilErrImpl := func() error {
		return &testErrorImpl{}
	}
	require.False(t, govalue.Nil(returnsNoneNilErrImpl()))

	err2 := returnsNoneNilErrImpl()
	require.False(t, govalue.Nil(err2))

	returnsNoneNilStruct := func() *testInterfaceImpl {
		return &testInterfaceImpl{}
	}
	require.False(t, govalue.Nil(returnsNoneNilStruct()))

	vs := returnsNoneNilStruct()
	require.False(t, govalue.Nil(vs))

	returnsStruct := func() testInterfaceImpl {
		return testInterfaceImpl{}
	}
	require.False(t, govalue.Nil(returnsStruct()))

	ns := returnsStruct()
	require.False(t, govalue.Nil(ns))
}

func TestNilAfterSet(t *testing.T) {
	err := fmt.Errorf("test error")
	err = nil
	require.True(t, govalue.Nil(err))

	ii := &testInterfaceImpl{}
	ii = nil
	require.True(t, govalue.Nil(ii))

	assertCallFuncWithInterface := func(t *testing.T, actual testInterface) {
		require.True(t, govalue.Nil(actual))
	}

	var v testInterface = &testInterfaceImpl{}
	v = nil
	require.True(t, govalue.Nil(v))
	assertCallFuncWithInterface(t, v)

	m := make(map[string]struct{})
	m = nil
	require.True(t, govalue.Nil(m))

	s := make([]string, 0, 0)
	s = nil
	require.True(t, govalue.Nil(s))

	f := func() {}
	f = nil
	require.True(t, govalue.Nil(f))

	ch := make(chan struct{})
	ch = nil
	require.True(t, govalue.Nil(ch))
}

func TestNotNilFunc(t *testing.T) {
	require.False(t, govalue.NotNil(nil))

	var err error
	require.False(t, govalue.NotNil(err))

	var f func()
	require.False(t, govalue.NotNil(f))

	var ii *testInterfaceImpl
	require.False(t, govalue.NotNil(ii))

	var ch chan struct{}
	require.False(t, govalue.NotNil(ch))

	var sl []string
	require.False(t, govalue.NotNil(sl))

	var ival *int
	require.False(t, govalue.NotNil(ival))

	var m map[string]struct{}
	require.False(t, govalue.NotNil(m))

	require.True(t, govalue.NotNil(0))
	require.True(t, govalue.NotNil(0.0))
	require.True(t, govalue.NotNil(""))
	require.True(t, govalue.NotNil(struct{}{}))
	require.True(t, govalue.NotNil(&testInterfaceImpl{}))
	require.True(t, govalue.NotNil(func() {}))
	require.True(t, govalue.NotNil(make(chan struct{})))
	require.True(t, govalue.NotNil(make([]string, 0)))
	require.True(t, govalue.NotNil(make(map[string]string)))
	require.True(t, govalue.NotNil(true))
	var i = 1
	ival = &i
	require.True(t, govalue.NotNil(ival))
}
