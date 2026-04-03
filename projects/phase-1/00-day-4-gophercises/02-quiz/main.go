package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func parseProblems(records [][]string) []problem {
	problems := make([]problem, 0, len(records))
	for _, record := range records {
		if len(record) != 2 {
			continue
		}
		problems = append(problems, problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		})
	}
	return problems
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file.")
		os.Exit(1)
	}

	problems := parseProblems(records)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		answerCh := make(chan string, 1)

		go func() {
			var answer string
			fmt.Scan(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTime's up! You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer == strings.TrimSpace(strings.ToLower(p.a)) {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
