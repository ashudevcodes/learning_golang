package main

import (
	"fmt"
)

func sum(num ...int) int {
	totle := 0
	for _, numbes := range num {
		totle += numbes
	}
	return totle
}

// interface id like any
func myPrint(value ...interface{}) interface{} {
	return value

}

func main() {
	result := sum(2, 5, 2, 1)
	nums := []int{1, 3, 53, 5}
	result1 := sum(nums...)
	fmt.Println(result1)
	fmt.Println(result)

}
