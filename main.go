package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", " a csv file in the format of question,answer")

	flag.Parse()
	
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file %s\n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse provided CSV file")
	}
	
	problems := parseLines(lines)
	correctAnswers := 0
	
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			fmt.Println("Correct!")
			correctAnswers++
		} else {
			fmt.Println("you are ngmi")
		}
	}
	fmt.Printf("you got %d correct out of %d\n", correctAnswers, len(problems))
}

type problem struct {
	question string
	answer string 
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return problems 
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}