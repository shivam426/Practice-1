package main

import (
	"fmt"
)

func main() {
	var str string
	var j int
	fmt.Println("enter the name:")
	fmt.Scan(&str)
	fmt.Println("enter the number:")
	fmt.Scan(&j)
	fmt.Println(str)
	fmt.Println(check(j))
}

func check(n int) bool {
	flag := false
	for i := 0; i <= 10; i++ {
		if i == n {
			flag = true
		}
	}
	return flag
}
