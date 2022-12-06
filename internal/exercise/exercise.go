package exercise

import (
	"bufio"
	"os"
)

type Exercise struct {
	Question   string
	Answer     string
	UserAnswer string
	Right      bool
}

func (e *Exercise) ScanAnswer(doneChan chan struct{}) {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	e.UserAnswer = scan.Text()
	e.checkUserAnswer()
	doneChan <- struct{}{}
}

func (e *Exercise) checkUserAnswer() {
	if e.UserAnswer == e.Answer {
		e.Right = true
	}
}

func (e *Exercise) IsAnswered() bool {
	return e.UserAnswer != ""
}
