package offshoot

import (
	"testing"
)

func TestOr_CreateNewOffshoot_WhenSuccess(t *testing.T) {

	// ARRANGE
	value := "hello"
	offshoot := Create[string](value)

	// ACT
	or := offshoot.Or(
		func(value string) Offshoot[string] {
			// Positive state must be handled
			return Create("mutated")
		},
		func(err error) Offshoot[string] {
			// Failed state must be handled
			return Create("fix error")
		})

	// ASSERT
	if or.Value() != "mutated" {
		t.Errorf("got %q, wanted %q", offshoot.Value(), value)
	}
}

func TestOr_CreateNewError_WhenFailure(t *testing.T) {

	// ARRANGE
	value := "hello"
	offshoot := Create[string](value)
	offshoot.Ensure(func(input string) bool {
		return false
	})

	// ACT
	or := offshoot.Or(
		func(value string) Offshoot[string] {
			// Positive state must be handled
			return Create("mutated")
		},
		func(err error) Offshoot[string] {
			// Failed state must be handled
			return Create("fix error")
		})

	// ASSERT
	if or.Value() != "fix error" {
		t.Errorf("got %q, wanted %q", offshoot.Value(), value)
	}
}
