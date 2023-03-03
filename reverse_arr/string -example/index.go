package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, a, b, c string
	a = "apple"
	b = "ball"
	c = "cat"
	s = "snake"
	fmt.Println(StringExample(a))
	fmt.Println(StringExample(b))
	fmt.Println(StringExample(c))
	fmt.Println(StringExample(s))
}

// StringExample for multiple string
func StringExample(str string) string {
	res := strings.ToUpper(str)
	return res
}
