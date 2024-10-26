package main

import (
	"fmt"
	"time"
)

func main() {
	i := 4
	switch i {
	case 3:
		fmt.Println("three")
	case 2:
		fmt.Println("Two")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Other")
	}

	// Multiple Conditon switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")

	}

	// Type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Is an integer")
		case string:
			fmt.Println("Its String", t)
		}
	}
	whoAmI("ashish")

}
