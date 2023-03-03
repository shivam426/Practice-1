package main

import "fmt"

func main() {

	var n, r int
	fmt.Scan(&r)
	for n = 1; n <= r; n++ {
		temp := n / 2
		count := 0
		for i := 2; i < temp; i++ {
			if n%i == 0 {
				count++
				break
			}

		}
		if count == 0 {
			fmt.Println(n)
		}
	}

}
