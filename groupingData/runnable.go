package main

import (
	s "example/hello/groupingData/slice"
	"fmt"
)

func main() {
	fmt.Println("---Arrays---")
	s.Arrays()
	fmt.Println("---Slice")
	s.Slices()
	fmt.Println("---For range slices---")
	s.ForRangeSlices()
	fmt.Println("---Adapting slices---")
	s.AdaptingSlices()
	fmt.Println("---Slice make---")
	s.SliceMake()
}
