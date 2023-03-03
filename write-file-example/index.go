package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("city.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// buf := make([]byte, 1024)
	// if _, err := file.Write(buf[:]); err != nil {
	// 	panic(err)
	// }
	// data := []string{"A", "b", "c", "d"}
	file.WriteString("`mumbai,` `Delhi,` `Bhubaneswar`")
}
