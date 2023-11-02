package main

import (
	"fmt"

	"github.com/ewilliams0305/offshoot"
)

func main() {
	fmt.Println("hello world")
	dowhat := offshoot.Create[string]("hello")

	fmt.Println(dowhat.IsSuccessfull())
	fmt.Println(dowhat.Error())
	fmt.Println(dowhat.Value())

	dowhat.Ensure(func(input string) bool {
		fmt.Println("Ensuring the data is valid")
		return input == "hello"
	})

	fmt.Println(dowhat.IsSuccessfull())

	dowhat.Ensure(func(input string) bool {
		fmt.Printf("Ensuring the data is valid %s \n", input)
		return false
	})

	fmt.Println(dowhat.IsSuccessfull())
}
