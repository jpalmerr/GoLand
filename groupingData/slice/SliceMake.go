package slice

import "fmt"

func SliceMake() {
	// si := []string{"a", "b", "c"}
	// fmt.Println(si)

	xi := make([]int, 0, 10) // essentially defines slice without populating it
	fmt.Println(xi)
	fmt.Println(len(xi)) // length
	fmt.Println(cap(xi)) // capacity
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("------------")
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi)) // 20 as append adds capacity of 10/same as slice its appending -> due to underlying array getting copied
}

//multidimensional slice

func MultiDimensionalSlice() {
	jb := []string{"James", "Bond", "Martini"}
	jm := []string{"Jenny", "Moneypenny", "Guinness"}

	xxs := [][]string{jb, jm}

	fmt.Println(xxs)

}
