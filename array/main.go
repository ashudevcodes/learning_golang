package main

import "fmt"

func main() {
	var num [4]bool
	fmt.Println(len(num))

	num[0] = true

	fmt.Println(num)

	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	arr := [2][2]int{{2, 3}, {4, 5}}
	fmt.Println(arr)

}
