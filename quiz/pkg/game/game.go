package game

import (
	"bufio"
	"fmt"
	"os"
)

type Quiz struct {
	question string
	answer   string
}

type Game struct {
	questions    []Quiz
	correctCount int
}

func (game *Game) Create(quizContents [][]string) {
	for _, value := range quizContents {

		if len(value) < 2 {
			panic("Incorrect slice of question and answers")
		}
		quiz := Quiz{question: value[0], answer: value[1]}

		game.questions = append(game.questions, quiz)
	}
}

func (game *Game) Start() {
	input := bufio.NewScanner(os.Stdin)
	for index, value := range game.questions {
		fmt.Println("Question No. ", index+1, value.question)
		input.Scan()
		answer := input.Text()
		if answer == value.answer {
			game.correctCount += 1
		}
	}
}

func (game *Game) PrintScore() {
	fmt.Println(game.correctCount, "/", len(game.questions))
}

func (game *Game) PrintQuiz() {
	for _, value := range game.questions {
		fmt.Println(value)
	}
}
