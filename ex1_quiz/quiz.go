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

func startQuiz(problems [][2]string) uint {
	var score uint
	reader := bufio.NewReader(os.Stdin)

	for _, qa := range problems {
		fmt.Print(qa[0], " = ")

		if input, _ := reader.ReadString('\n'); strings.Compare(strings.TrimSpace(input), qa[1]) == 0 {
			score++
		}
	}

	return score
}

func main() {
	problemsFilename := parseOptions("problems.csv", "The CSV file containing problems")["file"]
	problems := fetchProblems(problemsFilename)
	score := startQuiz(problems)

	fmt.Println("Your score is:", score)
}
