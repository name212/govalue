# govalue
Additional library for checking go interfaces on nil

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
	if govalue.IsNil(d) {
		err = &errorImpl{msg: "got nil do"}
	} else {
		d.Do()
	}

	return err
}

func main() {
	var d do
	err := callDo(d)
	if !govalue.IsNil(err) {
		fmt.Printf("call do got err: %s\n", err.Error())
	}

	d = &doImpl{}
	err = callDo(d)
	if !govalue.IsNil(err) {
		fmt.Printf("call do got err: %s\n", err.Error())
	}

	fmt.Printf("IsNil: %v\n", govalue.IsNil(nil))
	fmt.Printf("IsNil: %v\n", govalue.IsNil(false))
	fmt.Printf("IsNil: %v\n", govalue.IsNil(struct{}{}))

	var e error
	fmt.Printf("IsNil: %v\n", govalue.IsNil(e))
	fmt.Printf("IsNil: %v\n", govalue.IsNil(errors.New("not nil error")))

	var slice []string
	fmt.Printf("IsNil: %v\n", govalue.IsNil(slice))
	slice = append(make([]string, 0, 1), "hello")
	fmt.Printf("IsNil: %v\n", govalue.IsNil(slice))

	var m map[string]struct{}
	fmt.Printf("IsNil: %v\n", govalue.IsNil(m))
	m = make(map[string]struct{})
	fmt.Printf("IsNil: %v\n", govalue.IsNil(m))

	var f func()
	fmt.Printf("IsNil: %v\n", govalue.IsNil(f))
	f = func() {}
	fmt.Printf("IsNil: %v\n", govalue.IsNil(f))
}
```