package main

import (
	"fmt"

	"github.com/JayBee007/gophercises/quiz/pkg/csv"
)

func main() {
	records := csv.ReadCsvFile("internal/problems.csv")

	for index, outer := range records {
		question, answer := outer[0], outer[1]

		fmt.Println("No.", index, "Question:", question, "Answer:", answer)
	}
}
