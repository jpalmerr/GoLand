package slice

import (
	"fmt"
	"sort"
)

func SliceInternals() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")

	a[0] = 7

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")
	/**
	notice b changes also. This is because b is pointing at an underlying a array
	a [7 1 2 3 4 5]
	b [7 1 2 3 4 5]
	*/
}

func UnderlyingArrayCopy() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	copy(b, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")

	a[0] = 7

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("--------------")

	/**
	b stays the same - different underlying array
	a  [7 1 2 3 4 5]
	b  [0 1 2 3 4 5]
	*/
}

func SliceInternalsProgram() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)

	y := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(y))
	fmt.Println("after medianTwo", y)
}

// uses the same underlying array
// everything is pass by value in go =>
// the value is being passed into the func
// and then assigned to x
func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
	}
	return (n[i-1] + n[i]) / 2
}
