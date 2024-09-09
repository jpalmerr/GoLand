package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectStatements() {
	/**
	The make keyword in Go is used to create and initialize slices, maps, and channels.
	It allocates and initializes an object of type slice, maps, or channel,
	and returns a value of that type (not a pointer to it).
	*/

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	// go keyword launches a "go routine"
	/**
	A goroutine is a lightweight thread managed by the Go runtime.
	Goroutines are a fundamental part of Go’s concurrency model, allowing you to run functions concurrently.

	the go keyword is placed before a function call to execute that function in a new goroutine.
	This allows the program to perform multiple tasks simultaneously, which is useful for tasks that can be
	done independently or in parallel.
	*/
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41 // this sends 41 into ch1 channel
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}
}
