package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("file", "../csv/problems.csv", "Enter file name")
	gettime := flag.Int("time", 30, "Enter time limit")
	flag.Parse()

	openFile, err := os.Open(*filename)
	if err != nil {
		fmt.Println("No file found")
		return
	}
	lines, err := csv.NewReader(openFile).ReadAll()
	if err == io.EOF {
		fmt.Println("Reached end of the file")
		return
	}
	q := readFile(lines)

	correct := 0
	quiztime := time.NewTimer(time.Duration(*gettime) * time.Second)

	for index, questions := range q {
		select {
		case <-quiztime.C:
			fmt.Printf("Oops! you ran outta time.You scored %d out of %d\n", correct, len(q))
			return
		default:
			fmt.Printf("\nProblem %d -> %s = ", index+1, questions.ques)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if questions.ans == answer {
				correct++
			}
		} //select
	} //for
	fmt.Println("You scored %d out of %s\n", correct, len(q))
} //main

func readFile(lines [][]string) []quiz {
	q := make([]quiz, len(lines))
	for i, line := range lines {
		q[i] = quiz{
			ques: line[0],
			ans:  strings.TrimSpace(line[1]),
		}
	} //for
	return q
} //readFile

type quiz struct {
	ques string
	ans  string
}
