package main

import "fmt"

func main() {

	names := []string{"Alice", "John", "Mark"}
	fmt.Println(names)

	items := make([]int, 3, 5)

	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	items = append(items, 4)
	fmt.Printf("Items: %+v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	fmt.Printf("%+v", items[3:7])

}
