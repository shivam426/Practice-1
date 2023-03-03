package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	var i int
	var b []int
	var arr []int
	n := len(a)
	for i = 0; i < n; i++ {
		if a[i] == a[n-1] {

			arr = append(arr, a[i])
			a = a[:n-1]

		}
		b = append(arr, a...)
	}
	fmt.Println(b)
}
