package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int { // it can be x int y int
	return x + y
}

func sayHello() {
	fmt.Println("Hello")
}

/*
*
in Go return values may be named. If so they are treated as variables
a return statement without argument returns the named return values.
This is known as a "naked return"
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1 //bitwise left shift
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// constants - think scala val
const Pi = 3.142 // cannot be declared using :=
const PiTyped float64 = 3.142

const (
	// create a huge number by shifting a 1 bit left 100 places
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("random number", rand.Int())

	fmt.Printf("Now you have  %g problems \n", math.Sqrt(9))

	fmt.Println(add(42, 10))
	sayHello()

	fmt.Println(split(10))

	fmt.Println(swap("hello", "world"))

	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)

	var x, y int = 4, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// type inference
	var yy int = 77
	v := 7 // type inferred - it infers to the precision of the constant
	w := 0.142
	fmt.Printf("Type: %T Value: %v \n", yy, yy)
	fmt.Printf("Type: %T Value: %v \n", v, v)
	fmt.Printf("Type: %T Value: %v \n", w, w)

	fmt.Println(needInt(10))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
