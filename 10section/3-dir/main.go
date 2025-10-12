package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
	}

}
