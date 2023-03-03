package main

import (
	"fmt"
	// "sort"
)

func main() {
	arr := []int{1, -1, 3, 2, -7, -5, 11, 6}
	var i int
	var a []int
	n := len(arr)
	var b []int
	var c []int
	for i = 0; i < n; i++ {
		// sort.Ints(arr)
		if arr[i] > 0 {
			a = append(a, arr[i])
		}
		if arr[i] < 0 {
			b = append(b, arr[i])
		}
		c = append(a, b...)
	}

	fmt.Println(c)
	// fmt.Println(a)
}
