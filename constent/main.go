package main

import "fmt"

const age = 21

const (
	PORT = 3000
	host = "localhost"
)

func main() {
	const name = "Ashish"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(host, PORT)
}
