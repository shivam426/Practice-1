package main

import (
	"fmt"
)

func main() {
	var count int
	var sum,temp,sumsq int
	fmt.Printf("enter the count")
	fmt.Scanf("%d", &count)
	for i := 0; i <= count; i++ {
		square := i * i
		temp = square + temp
		sum = i + sum
		sumsq = sum * sum
	}
	fmt.Println( temp, sumsq)
}
