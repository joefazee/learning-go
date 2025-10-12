package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"Alice": 90,
		"James": 85,
		"Dan":   60,
	}
	fmt.Printf("%+v\n", studentGrades)
	studentGrades["Alice"] = 95
	fmt.Printf("%+v\n", studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Printf("Alice: %+v\n", alice)
	}

	key := "James"
	if value, ok := studentGrades[key]; ok {
		fmt.Printf("%s: %+v\n", key, value)
	}
	delete(studentGrades, "Alice")

	fmt.Printf("%+v\n", studentGrades)

	configs := make(map[string]int)
	fmt.Printf("%+v %T\n", configs, configs)

	if configs == nil {
		fmt.Printf("Config is nil")
	}

}
