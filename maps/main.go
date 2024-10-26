package main

import (
	"fmt"
)

// map -> hash, object, dict
func main() {

	m := make(map[string]string)
	m["name"] = "ashish"
	m["lastName"] = "prasad"
	fmt.Println(m["name"], m["lastName"])
	fmt.Println(m["phone"])
	delete(m, "name")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	ma := map[string]int{"Phone": 40, "Price": 23}
	fmt.Println(ma)
	// find element in map
	_, ok := ma["Phone"]
	if ok {
		fmt.Println("all ok")

	} else {
		fmt.Println("Not ok")
	}
}
