package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most use contruct in go
// + useful methods
func main() {
	// Uninitlized slices is null
	var nums []int
	fmt.Println(len(nums))

	var num = make([]int, 0, 5)
	// capacity -> max number of element can fit
	num = append(num, 1)
	// copy function
	var nums2 = make([]int, len(num))
	nums2 = append(nums2, 2)

	copy(nums2, num)

	fmt.Println(nums2)
	fmt.Println(cap(num))

	var nums3 = []int{1, 2, 4}

	fmt.Println(slices.Equal(nums3, nums2))

	fmt.Println(nums3[0:2])

}
