package main

import (
	"fmt"
)

func main() {
	a := []int{64, 34, 25, 12, 22, 11, 90}
	var i, j int
	n := len(a)
	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}
		}

	}
	fmt.Println(a)

}
