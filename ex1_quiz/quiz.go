package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func parseOptions(defaultFilename, description string) map[string]string {
	file := flag.String("file", defaultFilename, description)

	flag.Parse()

	return map[string]string{
		"file": *file,
	}
}

func fetchProblems(filename string) [][2]string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error occured while reading %s: %s", filename, err))
		panic(err)
	}

	reader := csv.NewReader(strings.NewReader(string(data)))
	problems := [][2]string{}

	for {
		entry, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error: ", err)
		}

		problems = append(problems, [2]string{entry[0], entry[1]})
	}

	return problems
}

func startQuiz(problems [][2]string, ch chan bool) {
	reader := bufio.NewReader(os.Stdin)

	for _, qa := range problems {
		fmt.Print(qa[0], " = ")

		if input, _ := reader.ReadString('\n'); strings.Compare(strings.TrimSpace(input), qa[1]) == 0 {
			ch <- true
		}
	}
}

func startTimedQuiz(problems [][2]string, timeout uint) uint {
	var score uint
	addScoreCh, timer := make(chan bool), time.NewTimer(time.Duration(timeout)*time.Second)

	defer timer.Stop()
	go startQuiz(problems, addScoreCh)

	for {
		select {
		case <-addScoreCh:
			score++
		case <-timer.C:
			fmt.Println("Time out!")
			return score
		}
	}
}

func main() {
	const timeout uint = 30 // seconds

	problemsFilename := parseOptions("problems.csv", "The CSV file containing problems")["file"]
	problems := fetchProblems(problemsFilename)
	score := startTimedQuiz(problems, timeout)

	fmt.Println("Total number of questions: ", len(problems))
	fmt.Println("Your score is:", score)
}
