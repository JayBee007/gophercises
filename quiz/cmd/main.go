package main

import (
	"flag"
	"fmt"

	"github.com/JayBee007/gophercises/quiz/pkg/csv"
	"github.com/JayBee007/gophercises/quiz/pkg/game"
)

func main() {

	var filePath string
	var timeLimit int

	flag.StringVar(&filePath, "csv", "internal/problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&timeLimit, "limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	fileContents := csv.ReadCsvFile(filePath)

	fmt.Println(fileContents)

	game.NewGame(fileContents, timeLimit)
}
