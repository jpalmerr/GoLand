package main

import (
	"fmt"
)

// see format verbs https://pkg.go.dev/fmt#Printf

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

	fmt.Printf("%s is %d years old \t and the type is %T and %T", name, age, name, age)

	fmt.Println(' ')

	fmt.Println(
		`utf8 is cool,
		2nd line`,
	)

	// exercise

	fmt.Println("hello world 🌍") // ctr cmd space
	rawString := `Hello, World!
This is an example of a raw string literal in Go.
It can span multiple lines and include "quotes" without needing to escape them.
Here is an emoji: 🌍`

	fmt.Println(rawString)
}
