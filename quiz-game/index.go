package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problems struct {
	Question string
	Answer   string
}

func main() {
	file, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	Prb := ParsedLines(lines)
	ticker := time.NewTimer(time.Duration(20) * time.Second)
	Correct := 0
	for i, v := range Prb {
		fmt.Printf("Question %d: %s=", i+1, v.Question)
		answerch := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerch <- ans
		}()
		select {
		case <-ticker.C:
			fmt.Printf("\n%d Correct out of %d\n", Correct, len(Prb))
			return
		case ans := <-answerch:
			if v.Answer == ans {
				Correct++
			}
		}
	}
	fmt.Printf("%d Correct out of %d\n", Correct, len(Prb))
}
func ParsedLines(lines [][]string) []Problems {
	sel := make([]Problems, len(lines))
	for i, value := range lines {
		sel[i] = Problems{
			Question: value[0],
			Answer:   strings.TrimSpace(value[1]),
		}
	}
	return sel
}
