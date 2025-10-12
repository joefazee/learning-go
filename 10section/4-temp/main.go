package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("Removing tempFile", tempFile.Name())
		_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.Write([]byte("Hello World\n"))
	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}

	// temp dir
	tempDir, err := os.MkdirTemp("", "my_app_logs")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Removing tempDir", tempDir)
		_ = os.Remove(tempDir)
	}()

}
