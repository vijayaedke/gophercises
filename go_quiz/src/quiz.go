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

//strcuture quiz
type quiz struct {
	ques string
	ans  string
}

func main() {
	//command line arguments for filepath and time limit of quiz
	filename := flag.String("file", "../csv/problems.csv", "Enter file name")
	gettime := flag.Int("time", 30, "Enter time limit")
	flag.Parse()

	//open specified file path
	openFile, err := os.Open(*filename)
	if err != nil {
		fmt.Println("No file found")
		return
	}
	//read file content
	lines, err := csv.NewReader(openFile).ReadAll()
	//check if reached end of the file
	if err == io.EOF {
		fmt.Println("Reached end of the file")
		return
	}

	//create a structure of quiz and read file contents classifying questions and answer
	q := readFile(lines)

	//initializa counter for correct answers and set timer for quiz
	correct := 0
	quiztime := time.NewTimer(time.Duration(*gettime) * time.Second)

	//parse through questions and check for correct answer
	for index, questions := range q {
		//using select since we need to check the values via timer channel
		//switch is used when we need to check case based on value
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

//function to read file contents
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
