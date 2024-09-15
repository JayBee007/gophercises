package main

import (
	"fmt"
	"os"

	"github.com/JayBee007/gophercises/quiz/pkg/csv"
	"github.com/JayBee007/gophercises/quiz/pkg/game"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic: ", r)
	}
}

func main() {

	args := os.Args
	argsLength := len(args)

	defer recoverFromPanic()

	program := args[0]

	var flag string

	if argsLength > 1 {
		flag = args[1]
	}

	switch flag {
	case "-h", "--help":
		fmt.Println("Usage of ", program, ":\n\n\t -csv string\n\t\ta csv file in the format of 'question,answer' (default \"problems.csv\")\n\n\t -limit int\n\t\tthe time limit for the quiz in seconds (default 30)")
	case "-csv":
		filepath := "internal/problems.csv"
		if len(args) == 3 {
			filepath = args[2]
		}
		fileContents := csv.ReadCsvFile(filepath)
		game := game.Game{}
		game.Create(fileContents)
		game.Start()
		game.PrintScore()

	default:
		fmt.Println("Unknown command flag")
	}
}
