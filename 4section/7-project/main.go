package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division       = "Division"
	divisionErrMsg = "division by zero is not allowed"
)

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ","),
		e.Message)
}

func Sum(numbers ...int) int {
	defer fmt.Println("Sum finished")

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionErrMsg,
		}
	}

	return a / b, nil
}

func main() {

	fmt.Println(Sum(1, 2, 3))

	value, err := SafeDivision(10, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("SafeDivision", value)

}
