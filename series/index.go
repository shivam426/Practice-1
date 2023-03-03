package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var fact float64 = 1
	num := 2
	var sum float64 = 1
	// //var mul float64
	var i int
	// i, j = 1, 1
	// for i <= n {

	//
	// 	fmt.Println(mul)
	// 	sum += int(mul)
	// 	i = i + 1
	// 	j = j + 2
	// 	// div := int(mul) / fact
	// 	// sum = sum + div
	// 	// fmt.Println(sum)
	for i = 1; i < n; i++ {
		fact = fact * float64(i)
		var mul = math.Pow(float64(num), float64(2*i-1)) / fact
		// fmt.Println(mul)
		sum += mul
		fmt.Println(sum)
	}

}
