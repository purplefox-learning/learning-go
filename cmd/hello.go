package main

import (
	"fmt"

	"github.com/purplefox81/learning-go/stringutil"
)

func main() {
	fmt.Printf(Hello())
	fmt.Println(ReverseHello())
}

// Hello returns a simple hello world message
func Hello() string {
	return "Hello, world"
}

// ReverseHello reverse a given hello world message
func ReverseHello() string {
	return stringutil.Reverse("!oG ,olleH")
}
