package main

import "fmt"

/**
This code requires Go 1.26 or later to run, as it uses the new built-in function `new` that can take an expression as an argument. The `new` function is used to create a pointer to a value of the specified type, and in this case, it allows us to create pointers to boolean and integer values directly from expressions.
*/

type Request struct {
	HasPost *bool
	Limit   *int
	Query   string
}

func main() {
	// HasPost *bool
	// true | false | nil

	// Before go 1.26
	hasPost := true
	req := Request{HasPost: &hasPost}
	fmt.Printf("Old Appraoch Request: %+v\n", req)

	// new(expression)
	req = Request{HasPost: new(true), Limit: new(2 + 2), Query: "name"}
	fmt.Printf("New Appraoch Request: %+v\n", req)

}
