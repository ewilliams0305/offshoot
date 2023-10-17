package main

import (
	"errors"
	"fmt"
)

// The dowhat type can only ever be a success or a result.
// When the dowhat is a success it will store a value of type T
// when the dowhat is a failure it will store an error value.
type DoWhat[T any] struct {
	value   T
	error   error
	success bool
	failure bool
}

type MapperFunction[TInput any, TOutput any] func(input TInput) TOutput

type EnsureFunction[TInput any] func(input TInput) bool

func (dowhat *DoWhat[T]) Value() T {
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

func Create[T any](value T) DoWhat[T] {
	if &value != nil {
		return DoWhat[T]{
			success: true,
			failure: false,
			error:   nil,
			value:   value}
	}
	return DoWhat[T]{
		success: false,
		failure: true,
		error:   errors.New("The value of the defined type is nil")}
}

func Failure[T any](error error) DoWhat[T] {
	return DoWhat[T]{
		success: false,
		failure: true,
		error:   error}
}

// func (dowhat *DoWhat[T]) Map(mapper Mapper[T, TOutput]) DoWhat[TOutput] {
// 	if dowhat.success {
// 		var value T = mapping(dowhat.value)
// 	}

// 	return DoWhat[T]{value: mapping(dowhat.value)}
// }

func (dowhat *DoWhat[T]) Ensure(ensure EnsureFunction[T]) DoWhat[T] {
	if dowhat.success {

		if ensure(dowhat.value) {
			return *dowhat
		}
		return Failure[T](errors.New("Failed to ensure the value was valid"))
	}
	return Failure[T](errors.New("Failed to ensure the value was valid"))
}

func main() {
	fmt.Println("hello world")
	dowhat := Create[string]("hello")

	fmt.Println(dowhat.IsSuccessfull())
	fmt.Println(dowhat.error)
	fmt.Println(dowhat.value)
	fmt.Println(dowhat.Value())

	dowhat.Ensure(func(input string) bool {
		fmt.Println("Ensuring the data is valid")
		return input == "hello"
	})

	fmt.Println(dowhat.success)

	dowhat.Ensure(func(input string) bool {
		fmt.Printf("Ensuring the data is valid %s \n", input)
		return false
	})

	fmt.Println(dowhat.IsSuccessfull())
}
