package main

import (
	"fmt"
	"sort"
)

func main() {
	// var arr [5]int
	arr := []int{2, 3, 6, 4, 5, 1}
	n := len(arr)
	var i int
	fmt.Println("enter the element to be picked")
	for i = 0; i < n; i++ {
		// fmt.Scan(&arr[i])
		sort.Ints(arr)

	}
	fmt.Println(arr[2])
	fmt.Println(arr)

}
