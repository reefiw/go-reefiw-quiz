package config

import (
	"encoding/csv"
	"flag"
	"fmt"
	"go-reefiw-quiz/internal/exercise"
	"os"
)

type flags struct {
	time       int
	pathToFile string
}

type AppConf struct {
	//All time in sec
	TookTime     int
	DecisionTime int
	Exercises    []exercise.Exercise
}

func NewConfig() (*AppConf, error) {
	flagsStorage := parseFlags()
	appConf := AppConf{TookTime: 0, DecisionTime: flagsStorage.time}
	var err error
	appConf.Exercises, err = parseExercises(flagsStorage.pathToFile)
	if err != nil {
		return nil, err
	}
	return &appConf, err
}

func parseFlags() *flags {
	flagsStorage := flags{}
	path := flag.String("f", "./internal/csv/problems.csv", "path to file")
	time := flag.Int("t", 10, "path to file")
	flag.Parse()
	flagsStorage.pathToFile = *path
	flagsStorage.time = *time
	return &flagsStorage
}

func parseExercises(pathToFile string) ([]exercise.Exercise, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	fmt.Println(records)
	var exercises []exercise.Exercise
	for i := 0; i < len(records); i++ {
		exercises = append(exercises, exercise.Exercise{Question: records[i][0], Answer: records[i][1], Right: false})
	}
	return exercises, nil
}
