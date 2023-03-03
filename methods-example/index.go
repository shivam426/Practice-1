package main

import (
	"fmt"
)

func main() {
	fmt.Println(Add(57, 47))
}
func Add(x, y int) (int, bool) {
	val := x + y

	f := Even(val)
	return val, f
}

func Even(z int) bool {
	flag := false

	if z%2 == 0 {
		flag = true

	}
	return flag
}
