package main

import (
	"fmt"
	"math"
)

func main() {

	// var mod, sum int
	// rev := 1
	num := 123
	num1 := num
	// temp := num
	count := 0
	for num > 0 {

		// 	mod = num % 10
		// 	rev = mod * rev
		num = num / 10
		count++
		// 	sum = sum + rev

	}
	fmt.Println(count)
	var sum float64
	for num1 > 0 {
		mod := num1 % 10
		//fmt.Println(mod)
		rev := float64(mod) * math.Pow(10, float64(count-1))
		count--
		//fmt.Println(rev)
		sum = sum + rev
		num1 = num1 / 10
	}
	fmt.Println(sum)
}
