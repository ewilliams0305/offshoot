package main

import (
	"errors"

	"github.com/ewilliams0305/offshoot/offshoot"
)

func main() {
	exampleWithSuccess()
	exampleWithError()
}

func exampleWithSuccess() {
	pass := offshoot.New(functionWithNillError())
	pass.Ensure(func(input string) bool { return true })

	println("%s", pass.Value())
}

func exampleWithError() {
	fail := offshoot.New(functionWithError())
	fail.Or(
		func(value string) offshoot.Offshoot[string] {
			print("this should NOT print as the function contains an error and the hapy path is not invoked.")
			return offshoot.Create(value)
		},
		func(err error) offshoot.Offshoot[string] {
			print("this will print as the error is not nil, this is you chance to handle it")
			return offshoot.Create(err.Error())
		})
}

func functionWithError() (string, error) {
	return "hello", errors.New("this is an error")
}

func functionWithNillError() (string, error) {
	return "hello", nil
}
