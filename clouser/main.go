package main

import "fmt"

func counter() func() int {
	var count = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	vale := counter()
	vale()
	vale()
	fmt.Println(vale())

}
