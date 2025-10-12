package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path1 := filepath.Join("C:", "Users", "Documents")
	fmt.Println(path1)

	path2 := filepath.Join("config", "app.yaml")
	fmt.Println(path2)

	fmt.Println(filepath.Base(path2))
	fmt.Println(filepath.Ext(path2))

	dirtyDir := "./users/./dir/../other_dir/./file.txt"
	fmt.Println(filepath.Clean(dirtyDir))

}
