package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	// calssic for loop
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	// range
	for j := range 10 {
		fmt.Println(j)
	}
}
