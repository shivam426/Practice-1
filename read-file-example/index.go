package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// r, err := ioutil.ReadFile("name.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // var b []string
	// b := string(r)
	// fmt.Println(b)

	r, err := readLines("name.txt")

	if err != nil {
		fmt.Println(err)

	}
	for i, line := range r {

		fmt.Println(i, line)
	}
}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
