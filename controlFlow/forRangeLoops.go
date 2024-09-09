package main

import "fmt"

func forRangeLoops() {
	// for range loop
	// data structures - slice
	xi := []int{42, 43, 44, 45, 46, 47} // array of ints
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop
	// data structures - maps. maps are not ordered
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a maps", k, v)
	}

	/**
	If you're looping over an array, slice, string, or maps, or reading from a channel, a range clause can manage the loop
	*/

	fmt.Println("--- array --- ")
	// array
	array := [5]int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("--- slice --- ")
	// slice
	/*
		Slices: Dynamic size, can grow and shrink as needed. e.g you can append
	*/
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("---string---")
	//string
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	fmt.Println("---maps---")
	// maps
	oldMap := map[string]int{"a": 1, "b": 2, "c": 3}
	newMap := make(map[string]int) // think of make as setting a variable with a type but not a value, like a description
	for key, value := range oldMap {
		fmt.Printf("Key: %s, Value: %d ", key, value)
		newMap[key] = value
	}
	fmt.Println(newMap)

	reverseMap := make(map[int]string)
	for key, value := range oldMap {
		reverseMap[value] = key
	}
	fmt.Println(reverseMap)

	// just get key
	for key := range oldMap {
		fmt.Println("Key:", key)
		// Here you can add any logic you need to process the key
	}
}
