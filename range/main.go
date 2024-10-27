package main

import "fmt"

// iteraing over data structure
func main() {
	nums := []int{2, 5, 2}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	nums1 := []int{1, 4, 4}
	sum := 0
	for _, num := range nums1 {
		sum += num
	}
	fmt.Println(sum)

	nums2 := []int{12, 31, 5}
	for index, numbers := range nums2 {
		fmt.Println(numbers, index)
	}
	nums3 := map[string]string{"Name": "Ashish", "LastName": "Prasad"}
	for key, value := range nums3 {
		fmt.Println(key, value)
	}
	// unicode point character
	for in, alpha := range "GoLang" {
		fmt.Println(in, alpha)
	}

}
