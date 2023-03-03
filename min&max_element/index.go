package main

import (
	"fmt"
)

func main() {
	arr := []int{23, 45, 32, 65, 9, 12}
	min := arr[0]
	max := min
	n := len(arr)
	var i int
	for i = 0; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}

	}
	fmt.Println(max, min)
}
