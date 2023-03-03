package main

import (
	"fmt"
)

func main() {
	a := []int{64, 34, 25, 12, 22, 11, 90}
	var i, j int
	n := len(a)
	for i = 0; i < n; i++ {
		for j = 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				swp := a[j]
				a[j] = a[j+1]
				a[j+1] = swp
			}
		}
	}
	fmt.Println(a)
}
