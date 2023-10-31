/*
The offshoot package is used as a common type to store either an error or an value.
An offshoot can be returned by a potentially destructive method to capture either the successful resulting value or error.
An offshoot instance can only ever be a failure or a success and can never store both vales and errors at the same time.
*/
package offshoot

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

func New[T any](value T, err error) Offshoot[T] {
	if err == nil {
		return Create(value)
	}
	return Failure[T](err)
}

func Create[T any](value T) Offshoot[T] {

	return Offshoot[T]{
		success: true,
		failure: false,
		error:   nil,
		value:   value}
}

func Failure[T any](error error) Offshoot[T] {
	return Offshoot[T]{
		success: false,
		failure: true,
		error:   error}
}
