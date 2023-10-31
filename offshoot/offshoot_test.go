package offshoot

import (
	"strconv"
	"testing"
)

func TestCreate_StoresValue_WhenValid(t *testing.T) {

	// ARRANGE
	value := "hello"
	offshoot := Create[string](value)

	// ACT

	// ASSERT
	if offshoot.Value() != value {
		t.Errorf("got %q, wanted %q", offshoot.Value(), value)
	}
}

func TestEnsure_SucceedsOffsoot_WhenPredicateIsTrue(t *testing.T) {

	// ARRANGE
	value := 10
	offshoot := Create(value)

	// ACT
	offshoot.Ensure(func(input int) bool {
		return input == value
	})

	// ASSERT
	if offshoot.IsFailure() {
		t.Errorf("got %q, wanted %q", offshoot.Value(), value)
	}
}

func TestEnsure_FailsOffshoot_WhenPredicateIsFalse(t *testing.T) {

	// ARRANGE
	value := 10
	offshoot := Create(value)

	// ACT
	offshoot.Ensure(func(input int) bool {
		return input == 12
	})

	// ASSERT
	if offshoot.IsSuccessfull() {
		t.Errorf("got %q, wanted %q", offshoot.Value(), value)
	}
}

func TestEnsure_DoesNotValidate_WhenOffshootIsFailure(t *testing.T) {

	// ARRANGE
	value := 10
	offshoot := Create(value)

	// ACT
	off := offshoot.
		Ensure(func(input int) bool {
			return input == 12
		}).
		Ensure(func(input int) bool {
			return true
		})

	result := MapOffshoot[int, string](off, func(input int) string {
		return strconv.Itoa(input)
	})

	// ASSERT
	if offshoot.IsSuccessfull() {
		t.Errorf("got %q, wanted %q", result.Value(), "10")
	}
}
