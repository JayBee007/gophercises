package game

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Quiz struct {
	question string
	answer   string
}

type Game struct {
	questions    []Quiz
	correctCount int
	timeLimit    int
}

func NewGame(quizContents [][]string, timeLimit int) {
	game := Game{}
	game.create(quizContents, timeLimit)
	game.start()
}

func (game *Game) create(quizContents [][]string, timeLimit int) {
	game.timeLimit = timeLimit

	for _, value := range quizContents {

		if len(value) < 2 {
			panic("Incorrect slice of question and answers")
		}
		quiz := Quiz{question: value[0], answer: value[1]}

		game.questions = append(game.questions, quiz)
	}
}

func (game *Game) start() {

	timer := time.NewTimer(time.Duration(game.timeLimit) * time.Second)

	input := bufio.NewScanner(os.Stdin)

	go func() {
		for index, value := range game.questions {
			fmt.Println("Question No. ", index+1, value.question)
			input.Scan()
			answer := input.Text()
			if answer == value.answer {
				game.correctCount += 1
			}
		}
	}()

	if _, ok := <-timer.C; ok {
		printTimesUp()
		game.printScore()

		return
	}

	game.printScore()

}

func printTimesUp() {
	fmt.Println("Times up")
}

func (game *Game) printScore() {
	fmt.Println(game.correctCount, "/", len(game.questions))
}
