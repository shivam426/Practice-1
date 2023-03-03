package main

import (
	"fmt"
)

func main() {
	var i, j, tmp int
	a := []int{1, 5, 3, 4, 3, 5, 6, 3}
	n := len(a)

	for i = 0; i < n; i++ {
		tmp = 0
		for j = i + 1; j < n; j++ {
			if a[i] == a[j] {
				tmp = tmp + 1
				fmt.Println(tmp)
			}

		}
	}
	fmt.Println(tmp)
}
