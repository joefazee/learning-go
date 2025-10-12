package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)

	b := strings.Builder{}
	b.WriteString("Here is an example")

	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	s3 := "     test sss    "
	fmt.Println("s3", len(s3))
	s3 = strings.TrimSpace(s3)

	fmt.Println("s3 after trim", len(s3))

	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))

	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jane example.com")
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ","))

}
