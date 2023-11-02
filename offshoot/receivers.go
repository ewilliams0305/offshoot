package offshoot

import (
	"errors"
)

// Ensure ensures that a certain condition is met for the value held by the Offshoot.
// If the Offshoot is successful and the condition is met, it returns the same Offshoot.
// If the condition is not met, it marks the Offshoot as failed and associates an error.
// If the Offshoot is already in a failed state, it remains unchanged.
//
// Parameters:
//   - ensure: A function that takes the value of the Offshoot and returns a boolean indicating
//     whether the value satisfies the condition.
//
// Returns:
//   - *Offshoot[T]: A pointer to the original Offshoot. If the condition is met, it returns
//     the same Offshoot; otherwise, it returns a failed Offshoot with an error.
//
// Example:
//
//	offshoot := Create("hello")
//	result := offshoot.Ensure(func(value T) bool {
//	  return value == "hello"
//	})
//	if result.Failed() {
//	  fmt.Println(result.Error())
//	}
func (offshoot *Offshoot[T]) Ensure(ensure EnsureFunction[T]) *Offshoot[T] {

	if offshoot.success {

		if ensure(offshoot.value) {
			return offshoot
		}
		return offshoot.fail(errors.New("FAILED TO ENSURE VALUE IS VALID"))
	}
	return offshoot
}

// Or returns either a value or an error depending on the success state of the Offshoot.
// If the Offshoot is successful, it invokes the 'value' function with the value held by the Offshoot
// and returns the result.
// If the Offshoot is in a failed state, it invokes the 'error' function with the error held by the Offshoot
// and returns the result.
//
// Parameters:
//   - value: A function that takes the value of the Offshoot and returns a new value.
//   - error: A function that takes the error of the Offshoot and returns a new error.
//
// Returns:
//   - Offshoot[T]: The result of invoking either the 'value' or 'error' function, depending on the Offshoot's state.
//
// Example:
//
//	offshoot := Create("Hello")
//	result := offshoot.Or(func(value T) T {
//	  return Create(value)
//	}, func(err error) error {
//	  // Your error-handling logic here
//	  return Failure(err)
//	})
func (offshoot *Offshoot[T]) Or(value ValueHandle[T], error ErrorHandle[T]) Offshoot[T] {
	if offshoot.success {
		return value(offshoot.value)
	}

	return error(offshoot.error)
}

// func (o *Offshoot[T]) toString() string {
// 	return fmt.Sprintf("RESULT: %+v ", o)
// }

// Mutates an existing offshoot with the error provided.
func (offshoot *Offshoot[T]) fail(e error) *Offshoot[T] {
	offshoot.error = e
	offshoot.failure = true
	offshoot.success = false
	return offshoot
}
