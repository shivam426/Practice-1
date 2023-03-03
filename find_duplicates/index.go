package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1, 2, 3}
	var i, j int
	n := len(a)
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if a[i] == a[j] {
				fmt.Println(a[i])
			} else {
				// fmt.Println("-1")
			}
		}
	}
}
