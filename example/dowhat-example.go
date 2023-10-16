package main

import (
	"errors"
	"fmt"
)

// The dowhat type can only ever be a success or a result.
// When the dowhat is a success it will store a value of type T
// when the dowhat is a failure it will store an error value.
type DoWhat[T any] struct {
	value   *T
	error   error
	success bool
	failure bool
}

func (dowhat *DoWhat[T]) Value() *T {
	return dowhat.value
}

func (dowhat *DoWhat[T]) Error() error {
	return dowhat.error
}

func (dowhat *DoWhat[T]) IsSuccessfull() bool {
	return dowhat.success
}

func (dowhat *DoWhat[T]) IsFailure() bool {
	return dowhat.failure
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func Create[T any](value *T) DoWhat[T] {
	if value != nil {
		return DoWhat[T]{
			success: true,
			failure: false,
			error:   nil,
			value:   value}
	}
	return DoWhat[T]{
		success: false,
		failure: false,
		error:   errors.New("The value of the defined type is nil"),
		value:   nil}
}

// func (dowhat *DoWhat[T]) Map(mapping func(input T) T) DoWhat[T] {
// 	if dowhat.success {
// 		var value T = mapping(dowhat.value)
// 	}

// 	return DoWhat[T]{value: mapping(dowhat.value)}
// }

func main() {
	fmt.Println("hello world")

	var str string = "122fefe"
	dowhat := Create[string](&str)

	fmt.Println(dowhat.IsSuccessfull())
	fmt.Println(dowhat.error)
	fmt.Println(dowhat.value)
	fmt.Println(dowhat.Value())

}
