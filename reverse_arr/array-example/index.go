package main

import (
	"fmt"
)

func main() {
	arr := [10]int{}
	a := [5]string{}
	fmt.Println("enter the numbers")
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("enter the name")
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])

	}

	fmt.Println(arr)
	fmt.Println(a)
}
