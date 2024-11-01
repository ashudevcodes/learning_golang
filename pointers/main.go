package main

import "fmt"

// By Value as Copy
func changeNum(num *int) {
	*num = 5
	fmt.Println("In chnageNum", *num)
}

func main() {
	num := 1
	// Pointers
	changeNum(&num)
	fmt.Println("Ater chnageNum in main", num)

}
