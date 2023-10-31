/*
The offshoot package is used as a common type to store either an error or an value.
An offshoot can be returned by a potentially destructive method to capture either the successful resulting value or error.
An offshoot instance can only ever be a failure or a success and can never store both vales and errors at the same time.
*/
package offshoot

import (
	"errors"
)

// The offshoot type can only ever be a success or a result.
// When the offshoot is a success it will store a value of type T
// when the offshoot is a failure it will store an error value.
type Offshoot[T any] struct {
	value   T
	error   error
	success bool
	failure bool
}

func (offshoot *Offshoot[T]) Value() T {
	return offshoot.value
}

func (offshoot *Offshoot[T]) Error() error {
	return offshoot.error
}

func (offshoot *Offshoot[T]) IsSuccessfull() bool {
	return offshoot.success
}

func (offshoot *Offshoot[T]) IsFailure() bool {
	return offshoot.failure
}

// // errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// func New(text string) error {
// 	return &errorString{text}
// }

func New[T any](value T, error Error) Offshoot[T] {
	if error = nil {
		return Create(value)
	}
	return Failure(error)
}

func Create[T any](value T) Offshoot[T] {
	if &value != nil {
		return Offshoot[T]{
			success: true,
			failure: false,
			error:   nil,
			value:   value}
	}
	return Offshoot[T]{
		success: false,
		failure: true,
		error:   errors.New("The value of the defined type is nil")}
}

func Failure[T any](error error) Offshoot[T] {
	return Offshoot[T]{
		success: false,
		failure: true,
		error:   error}
}

func (offshoot *Offshoot[T]) fail(e error) *Offshoot[T] {
	offshoot.error = e
	offshoot.failure = true
	offshoot.success = false
	return offshoot
}

// func (offshoot *Offshoot[TInput]) Mapper(mapper MapperFunction[TInput, TOutput]) Offshoot[TOutput] {

// 	if offshoot.success {
// 		return Create(mapper(offshoot.value))
// 	}

// 	return Create(offshoot.value)
// }

func MapperT[TInput any, TOutput any](offshoot *Offshoot[TInput], mapper MapperFunction[TInput, TOutput]) Offshoot[TOutput] {

	if offshoot.success {
		return Create(mapper(offshoot.value))
	}

	return Failure[TOutput](errors.New("Failed to map values, the offshoot was already in a failed state"))
}

func (offshoot *Offshoot[T]) Ensure(ensure EnsureFunction[T]) *Offshoot[T] {

	if offshoot.success {

		if ensure(offshoot.value) {
			return offshoot
		}
		return offshoot.fail(errors.New("Failed to ensure the value matches the predicate"))
	}
	return offshoot
}
