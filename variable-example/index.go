package main

import (
	"fmt"
)

func main() {
	// var dob string
	a := []int{20, 30, 67, 40, 23}
	sum := 0
	// fmt.Scan(&dob)
	// fmt.Println(CalculateYear(dob))
	for i := 0; i < 5; i++ {
		// fmt.Scan(&a[i])
		sum = sum + a[i]
	}
	avg := float64(sum / 5)
	fmt.Println(avg)
}

// CalculateYear to calculate the age
// func CalculateYear(n string) string {
// 	var age string
// 	t := time.Now().Format("2006-01-02")
// 	cur, err := time.Parse(t, n)
// 	if err != nil {
// 		panic(err)
// 	}

// 	age = t.Sub(cur)

// 	return age
// }
