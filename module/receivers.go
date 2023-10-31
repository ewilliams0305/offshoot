package offshoot

import "errors"

func (offshoot *Offshoot[T]) Ensure(ensure EnsureFunction[T]) *Offshoot[T] {

	if offshoot.success {

		if ensure(offshoot.value) {
			return offshoot
		}
		return offshoot.fail(errors.New("FAILED TO ENSURE VALUE IS VALID"))
	}
	return offshoot
}

func (offshoot *Offshoot[T]) Or(value ValueHandle[T], error ErrorHandle[T]) Offshoot[T] {
	if offshoot.success {
		return value(offshoot.value)
	}

	return error(offshoot.error)
}
