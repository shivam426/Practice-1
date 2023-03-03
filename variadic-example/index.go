package main

import (
	"fmt"
)

func sum(names ...string) {

	for _, num := range names {
		fmt.Println(num)
	}
}

func main() {
	sum("shiv", "abhijit", "aman")
}
