package main

import "fmt"

func main() {
	if age := 21; age >= 18 {
		fmt.Println("Adult")
	}

	name := "ashish"
	if name == "ashish" {
		fmt.Println("Hello Ashish")
	} else if name == "prince" {
		fmt.Println("Hello Prince")
	} else {
		fmt.Println("Hello")

	}

}
