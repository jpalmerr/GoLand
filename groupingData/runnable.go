package main

import (
	m "example/hello/groupingData/maps"
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
	fmt.Println("---Slice multidimensional---")
	s.MultiDimensionalSlice()
	fmt.Println("---Slice internals---")
	s.SliceInternals()
	fmt.Println("---Slice internals: copy---")
	s.UnderlyingArrayCopy()
	fmt.Println("---Slice internals eg---")
	s.SliceInternalsProgram()

	fmt.Println("---Maps Intro---")
	m.MapsIntro()
}
