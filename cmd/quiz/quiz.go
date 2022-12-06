package main

import (
	"context"
	"fmt"
	"go-reefiw-quiz/internal/config"
	"go-reefiw-quiz/internal/statistic"
	"os"
	"time"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Quiz started! Answer questions below. Decision time: %d Good luck :)\n", conf.DecisionTime)
	gameStatistic := statistic.Statistic{}

	gameStatistic.Before()
	startQuiz(conf)
	gameStatistic.After()

	fmt.Println("Quiz finish!")
	gameStatistic.UpdateStatistic(conf.Exercises)
	gameStatistic.PrintStatistic()
}

func startQuiz(conf *config.AppConf) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.DecisionTime)*time.Second)
	defer cancel()
	for i := 0; i < len(conf.Exercises); i++ {
		exercise := &conf.Exercises[i]
		fmt.Printf("Question %d. %s is ?\n Your answer: ", i, exercise.Question)
		scanCh := make(chan struct{})
		go exercise.ScanAnswer(scanCh)
		select {
		case <-scanCh:
		case <-ctx.Done():
			return
		}
	}
}
