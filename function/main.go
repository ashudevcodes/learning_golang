package main

import "fmt"

func add(a int, b int) int {
	return a + b

}

// a anb d are going to be int data type we dont define saperatilt
func add1(a, d int) int {
	return a + d
}

// Multiple vale return feature
func getLanguage() (string, string, string) {
	return "GoLang", "C++", "TypeScript"

}

func processIt(fn func(a, d int) int) {
	fn(3, 3)
}

func processIt1() func(a int) int {
	return func(a int) int {
		return 2

	}
}

func main() {

	result := add(2, 5)
	fmt.Println(result)
	lang1, lang2, lang3 := getLanguage()
	fmt.Println(lang2, lang3, lang1)
	in := processIt1()
	in(3)
}
