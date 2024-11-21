package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("csv", "problem.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	// _ = fileName

	file, err := os.Open(*fileName)
	if err != nil {
		// panic(err) // improve this later
		failMsg := fmt.Sprintf("Failed to open the csv file: %s\n", *fileName)
		exit(failMsg)
	}
	// fmt.Println("file opened successfully")
	// _ = file

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		// panic(err) // improve this later
		exit("Failed to parse the provided CSV file.")
	}
	// fmt.Println(lines)

	problems := parseLines(lines)

	counter := 0

	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, prob.que)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == prob.ans {
			// fmt.Println("Correct")
			counter++
		}
	}
	fmt.Printf("You have scored %d out of %d\n", counter, len(problems))

}

// [[10+1 11] [7+3 10] [2+3 5]]
//    0    1

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			que: line[0],
			ans: line[1],
		}
	}
	return result
}

type problem struct {
	que string
	ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
