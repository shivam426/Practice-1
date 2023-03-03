package main

import (
	"fmt"
)

func main() {
	var i int
	a := []int{1, 2, 3, 4, 5, 6}
	var arr []int
	n := len(a)
	for i = n - 1; i >= 0; i-- {
		arr = append(arr, a[i])
	}
	fmt.Println(arr)
}
