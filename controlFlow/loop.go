package main

import (
	"fmt"
)

func loop() {
	//SEQUENCE
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// LOOPS / INTERATIVE
	// for statements

	/*
		It unifies for and while and there is no do-while.
		There are three forms, only one of which has semicolons.

		for init; condition; post { }

		// Like while
		for condition { }

		for { }

	*/
	// https://go.dev/doc/effective_go#for

	// ++ means add 1 to variable
	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v \t first  for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third  for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue // ie if this case (not even) continue to next iteration rather than break loop but don't need an output
		}
		fmt.Println("counting even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}
}
