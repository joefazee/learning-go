package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Monday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend! No work")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.23)

}
