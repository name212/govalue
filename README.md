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

type doAnother interface {
	DoAnother()
}

type doAnotherImpl struct{}

func (*doAnotherImpl) DoAnother() {}

type errorImpl struct {
	msg string
}

func (e *errorImpl) Error() string {
	return fmt.Sprintf("have error: %s", e.msg)
}

func callDo(d do) error {
	var err *errorImpl
	if govalue.IsNil(d) {
		err = &errorImpl{msg: "got nil do"}
	} else {
		d.Do()
	}

	return err
}

func isNilExample() {
	fmt.Printf("isNilExample\n\n")
	var d do
	err := callDo(d)
	if !govalue.IsNil(err) {
		fmt.Printf("call do empty IsNil check got err: %s\n", err.Error())
	}
	if govalue.IsNotNil(err) {
		fmt.Printf("call do empty IsNotNil check got err: %s\n", err.Error())
	}

	d = &doImpl{}
	err = callDo(d)
	if !govalue.IsNil(err) {
		fmt.Printf("call do not empty IsNil check got err: %s\n", err.Error())
	}
	if govalue.IsNotNil(err) {
		fmt.Printf("call do not empty IsNotNil check got err: %s\n", err.Error())
	}

	fmt.Printf("nil IsNil: %v\n", govalue.IsNil(nil))
	fmt.Printf("false IsNil: %v\n", govalue.IsNil(false))
	fmt.Printf("struct{} IsNil: %v\n", govalue.IsNil(struct{}{}))

	fmt.Printf("nil IsNotNil: %v\n", govalue.IsNotNil(nil))
	fmt.Printf("false IsNotNil: %v\n", govalue.IsNotNil(false))
	fmt.Printf("struct{} IsNotNil: %v\n", govalue.IsNotNil(struct{}{}))

	var e error
	fmt.Printf("nil error IsNil: %v\n", govalue.IsNil(e))
	fmt.Printf("error IsNil: %v\n", govalue.IsNil(errors.New("not nil error")))

	fmt.Printf("nil error IsNotNil: %v\n", govalue.IsNotNil(e))
	fmt.Printf("error IsNotNil: %v\n", govalue.IsNotNil(errors.New("not nil error")))

	var slice []string
	fmt.Printf("nil slice IsNil: %v\n", govalue.IsNil(slice))
	fmt.Printf("nil slice IsNotNil: %v\n", govalue.IsNotNil(slice))
	slice = append(make([]string, 0, 1), "hello")
	fmt.Printf("not nil slice IsNil: %v\n", govalue.IsNil(slice))
	fmt.Printf("not nil slice IsNotNil: %v\n", govalue.IsNotNil(slice))

	var m map[string]struct{}
	fmt.Printf("nil map IsNil: %v\n", govalue.IsNil(m))
	fmt.Printf("nil map IsNotNil: %v\n", govalue.IsNotNil(m))
	m = make(map[string]struct{})
	fmt.Printf("not nil map IsNil: %v\n", govalue.IsNil(m))
	fmt.Printf("not nil map IsNotNil: %v\n", govalue.IsNotNil(m))

	var f func()
	fmt.Printf("nil func IsNil: %v\n", govalue.IsNil(f))
	fmt.Printf("nil func IsNotNil: %v\n", govalue.IsNotNil(f))
	f = func() {}
	fmt.Printf("not nil func IsNil: %v\n", govalue.IsNil(f))
	fmt.Printf("not nil func IsNotNil: %v\n", govalue.IsNotNil(f))

	var ch chan struct{}
	fmt.Printf("nil chan IsNil: %v\n", govalue.IsNil(ch))
	fmt.Printf("nil chan IsNotNil: %v\n", govalue.IsNotNil(ch))
	ch = make(chan struct{})
	fmt.Printf("not nil IsNotNil: %v\n", govalue.IsNil(ch))
	fmt.Printf("not nil IsNotNil: %v\n", govalue.IsNotNil(ch))
}

func isConvertToExample() {
	fmt.Printf("\n\nisConvertToExample\n\n")

	var d do = &doImpl{}

	res, err := govalue.ConvertTo[do, *doImpl](d)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("res - type: %t val: %v\n", res, res)
	err = callDo(res)
	fmt.Printf("err: %v\n", err)

	a := &doAnotherImpl{}
	res2, err := govalue.ConvertTo[doAnother, *doImpl](a)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("res2 - type: %t val: %v\n", res2, res2)
}

func main() {
	isNilExample()
	isConvertToExample()
}

```