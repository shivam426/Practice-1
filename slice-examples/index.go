package main

import (
	"fmt"
)

func main() {
	arr := [2]string{"Hello", "World"}
	s1 := arr[:1]
	fmt.Println(s1)
	s2 := arr[1:]
	fmt.Println(s2)
}
