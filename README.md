# govalue
Additional library for checking go interfaces on nil and another types routines

## Usage

```go
package main

import (
	"errors"
	"fmt"

	"github.com/name212/govalue"
)

type do interface {
	Do()
}

type doImpl struct{}

func (*doImpl) Do() {
	fmt.Println("Call Do on doImpl")
}

type errorImpl struct {
	msg string
}

func (e *errorImpl) Error() string {
	return fmt.Sprintf("have error: %s", e.msg)
}

func callDo(d do) error {
	var err *errorImpl
	if govalue.Nil(d) {
		err = &errorImpl{msg: "got nil do"}
	} else {
		d.Do()
	}

	return err
}

func NilExample() {
	fmt.Printf("NilExample\n\n")
	var d do
	err := callDo(d)
	if !govalue.Nil(err) {
		fmt.Printf("call do empty Nil check got err: %s\n", err.Error())
	}
	if govalue.NotNil(err) {
		fmt.Printf("call do empty NotNil check got err: %s\n", err.Error())
	}

	d = &doImpl{}
	err = callDo(d)
	if !govalue.Nil(err) {
		fmt.Printf("call do not empty Nil check got err: %s\n", err.Error())
	}
	if govalue.NotNil(err) {
		fmt.Printf("call do not empty NotNil check got err: %s\n", err.Error())
	}

	fmt.Printf("nil Nil: %v\n", govalue.Nil(nil))
	fmt.Printf("false Nil: %v\n", govalue.Nil(false))
	fmt.Printf("struct{} Nil: %v\n", govalue.Nil(struct{}{}))

	fmt.Printf("nil NotNil: %v\n", govalue.NotNil(nil))
	fmt.Printf("false NotNil: %v\n", govalue.NotNil(false))
	fmt.Printf("struct{} NotNil: %v\n", govalue.NotNil(struct{}{}))

	var e error
	fmt.Printf("nil error Nil: %v\n", govalue.Nil(e))
	fmt.Printf("error Nil: %v\n", govalue.Nil(errors.New("not nil error")))

	fmt.Printf("nil error NotNil: %v\n", govalue.NotNil(e))
	fmt.Printf("error NotNil: %v\n", govalue.NotNil(errors.New("not nil error")))

	var slice []string
	fmt.Printf("nil slice Nil: %v\n", govalue.Nil(slice))
	fmt.Printf("nil slice NotNil: %v\n", govalue.NotNil(slice))
	slice = append(make([]string, 0, 1), "hello")
	fmt.Printf("not nil slice Nil: %v\n", govalue.Nil(slice))
	fmt.Printf("not nil slice NotNil: %v\n", govalue.NotNil(slice))

	var m map[string]struct{}
	fmt.Printf("nil map Nil: %v\n", govalue.Nil(m))
	fmt.Printf("nil map NotNil: %v\n", govalue.NotNil(m))
	m = make(map[string]struct{})
	fmt.Printf("not nil map Nil: %v\n", govalue.Nil(m))
	fmt.Printf("not nil map NotNil: %v\n", govalue.NotNil(m))

	var f func()
	fmt.Printf("nil func Nil: %v\n", govalue.Nil(f))
	fmt.Printf("nil func NotNil: %v\n", govalue.NotNil(f))
	f = func() {}
	fmt.Printf("not nil func Nil: %v\n", govalue.Nil(f))
	fmt.Printf("not nil func NotNil: %v\n", govalue.NotNil(f))

	var ch chan struct{}
	fmt.Printf("nil chan Nil: %v\n", govalue.Nil(ch))
	fmt.Printf("nil chan NotNil: %v\n", govalue.NotNil(ch))
	ch = make(chan struct{})
	fmt.Printf("not nil NotNil: %v\n", govalue.Nil(ch))
	fmt.Printf("not nil NotNil: %v\n", govalue.NotNil(ch))
}

func main() {
	NilExample()
}

```