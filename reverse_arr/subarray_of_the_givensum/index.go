package main

import (
	"fmt"
)

func main() {
	var i, j, sum, start, end int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(a)
	for i = 0; i < n; i++ {
		for j = 1 + i; j < n; j++ {
			sum = (a[i] + a[j])
			if sum == 15 {
				start = 15 - a[i]
				end = 15 - a[j]
			}
		}

	}
	fmt.Println(start, end)
}
