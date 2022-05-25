package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	greetings := "Hello, OTUS!"
	greetings = stringutil.Reverse(greetings)
	fmt.Println(greetings)
}
