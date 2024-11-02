package main

import (
	"fmt"
)

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// also i can use comparable besides of generic type comparable means ony those
// data type the have compare to each other like int to int

func printGenericUsingComparable[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item)
		fmt.Println(name)
	}
}
func printGenericSlice[T string | float32](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T any] struct {
	element []T
}

func main() {
	myStack := Stack[int]{
		element: []int{1, 3, 5, 6},
	}
	myStack2 := Stack[string]{
		element: []string{"a", "s", "h", "i", "s", "h"},
	}
	fmt.Println(myStack2.element)
	fmt.Println(myStack.element)
	printSlice([]int{1, 24, 5, 6})
	printGenericSlice([]string{"Ashish", "Prasad"})
	printGenericSlice([]float32{23.4, 23.5, 21.5})

}
