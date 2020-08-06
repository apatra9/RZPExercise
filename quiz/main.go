package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type question struct {
	q string
	a string
}

func main() {

	file1 := flag.String("csv", "problems.csv", "Opening quiz data.")

	timelimit := flag.Int("limit", 10, "eelwj")
	flag.Parse()

	//_ = file1
	file, err := os.Open(*file1)
	if err != nil {
		fmt.Println("Error opening:", *file1, err)
		return
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll() //readsrecords
	if err != nil {
		fmt.Println("Error parsing:", *file1, err)
		return

	}

	//fmt.Println(lines)
	prob := makeproblems(lines)

	//fmt.Println(makeproblems(lines))
	correct := 0
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	for i, p := range prob {
		fmt.Println("Question number", i, "  is :  ", p.q, "=?")
		anschan := make(chan string)
		go func() {
			var ans string
			fmt.Println("Enter ur answer:")
			fmt.Scanf("%s", &ans)
			anschan <- ans
		}()

		select {

		case time1 := <-timer.C:
			fmt.Println("Time out!")
			fmt.Println("No. of correct ans:", correct)
			fmt.Println("No. of incorrect ans:", len(prob)-correct)
			_ = time1
			return

		case ans := <-anschan:

			//fmt.Println("Answern number", i, "   :   ", p.a)

			if p.a == ans {
				fmt.Println("Correct!")
				correct++
			}

		}
		//fmt.Println("No. of correct ans:", correct)
		//fmt.Println("No. of incorrect ans:", len(prob)-correct)
	}
}

func makeproblems(lines [][]string) []question {

	set := make([]question, len(lines))
	for i, line := range lines {
		set[i] = question{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return set
}
