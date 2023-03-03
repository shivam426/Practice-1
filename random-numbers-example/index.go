package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		rand.Seed(time.Now().UnixNano())
		randomNum := random(0, 6)
		fmt.Printf("Random number: %d\n", randomNum)
	}
}
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
