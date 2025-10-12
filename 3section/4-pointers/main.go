package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Println("val is nil")
		return
	}
	*val = *val * 10 // dereferencing
	fmt.Printf("modifyPointer: %+v\n", val)
}

func main() {

	num := 10
	modifyValue(num)
	fmt.Println(num)

	modifyPointer(&num)
	fmt.Println(num)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr grade: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", *(&gradePtr))

}
