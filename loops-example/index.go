package main

import (
	"fmt"
)

func main() {
	var n int
	sum := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
