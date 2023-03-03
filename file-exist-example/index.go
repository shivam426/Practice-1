package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("../prime_number.go/index.go"); err == nil {
		fmt.Println("file exist")
	} else {
		fmt.Println("file do not exist")
	}
}
