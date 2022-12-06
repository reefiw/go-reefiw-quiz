package statistic

import (
	"fmt"
	"go-reefiw-quiz/internal/exercise"
	"time"
)

type Statistic struct {
	ExerciseNum        int
	CorrectAnswerNum   int
	IncorrectAnswerNum int
	TookTime           int
	before             time.Time
	after              time.Time
}

func (s *Statistic) UpdateStatistic(exercises []exercise.Exercise) {
	s.ExerciseNum = len(exercises)
	s.CorrectAnswerNum = 0
	s.IncorrectAnswerNum = 0
	for _, curExercise := range exercises {
		switch {
		case curExercise.Right:
			s.CorrectAnswerNum++
		case !curExercise.Right && curExercise.IsAnswered():
			s.IncorrectAnswerNum++
		}
	}
}

func (s *Statistic) PrintStatistic() {
	fmt.Println("\n---STATISTIC---")
	fmt.Printf("Num of questions: %d\n", s.ExerciseNum)
	fmt.Printf("Corrected answers: %d\n", s.CorrectAnswerNum)
	fmt.Printf("Incorrected answers: %d\n", s.IncorrectAnswerNum)
	fmt.Printf("Took time: %d secs", s.TookTime)
}

func (s *Statistic) Before() {
	s.before = time.Now()
}

func (s *Statistic) After() {
	s.after = time.Now()
	s.TookTime = int(s.after.Sub(s.before).Seconds())
}
