package main

import (
	"fmt"
)

func main() {

	for i := 5; i >= 1; i-- {
		for k := 0; k <= 5; k++ {
			fmt.Print(" ")
		}
		for k := 0; k <= 5-i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := 0; i <= 5; i++ {
		for j := 5 - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}

}
