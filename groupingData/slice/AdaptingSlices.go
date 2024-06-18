package slice

import "fmt"

func AdaptingSlices() {
	fmt.Println("appending a slice")
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("-------------")
	// variadic parameter -> a function parameter that allows you to pass a variable number of arguments: a bit like * in params in scala
	xi = append(xi, 45, 46, 47, 99, 777)
	fmt.Println(xi)
	fmt.Println("-------------")

	fmt.Println("Slicing a slice")
	xi2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %#v\n", xi2)
	fmt.Println("-------------")

	// [inclusive:exclusive]
	fmt.Printf("xi2 - %#v\n", xi2[2:6]) // []int{2, 3, 4, 5}
	fmt.Println("-------------")

	// [:exclusive]
	fmt.Printf("xi2 - %#v\n", xi2[:7])
	fmt.Println("-------------")

	// [inclusive:]
	fmt.Printf("xi2 - %#v\n", xi2[4:])
	fmt.Println("-------------")

	// [:]
	fmt.Printf("xi2 - %#v\n", xi2[:]) // : give me everything
	fmt.Println("-------------")

	fmt.Println("Deleting from a slice")
	xi3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi3 - %#v\n", xi3)
	fmt.Println("-------------")

	xi3 = append(xi3[:4], xi3[5:]...) // gets rid of 4 and aooends to xi3
	fmt.Printf("xi3 - %#v\n", xi3)
	fmt.Println("-------------")
}
