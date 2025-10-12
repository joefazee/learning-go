package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	data := "Welcome to the wonderful world of Go!"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)

	encodedStr := "V2VsY29tZSB0byB0aGUgd29uZGVyZnVsIHdvcmxkIG9mIEdvIQ=="

	decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		log.Fatal(err)
	}

	if string(decodedStr) != data {
		log.Fatalf("decoded string does not match encoded data")
	}

	rawData := []byte{0xDE, 0xAD, 0xEF, 0xCA, 0xFE}
	binaryCodedToString := base64.StdEncoding.EncodeToString(rawData)

	fmt.Println(string(binaryCodedToString))

	b64Str := "3q3vyv4="
	decodedStr, err = base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.Fatal(err)
	}

}
