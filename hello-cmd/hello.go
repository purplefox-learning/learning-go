package main

import (
	"fmt"

	"github.com/purplefox81/learning-go/stringutil"
)

const (
	helloPrefix = "Hello, "
)

func main() {
	fmt.Printf(Hello())
	fmt.Println(ReverseHello())
}

// Hello returns a simple hello world message
func Hello() string {
	return "Hello, Go!"
}

// HelloString returns a Hello plus the given string
func HelloString(s string) string {
	return helloPrefix + s
}

// ReverseHello reverse a given hello world message
func ReverseHello() string {
	return stringutil.Reverse("!oG ,olleH")
}
