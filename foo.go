package main

import (
	"fmt"

	"github.com/amitkgupta/goodlearn/classifier"
	"github.com/amitkgupta/goodlearn/csvparse"
)

func main() {
	trainingSet, err := csvparse.DatasetFromPath("../golearnbenchmarks/classifiers/datasets/basic_training.csv", 4, 4)
	if err != nil {
		panic(err)
	}

	classifier, err := classifier.NewKNNClassifier(3)
	if err != nil {
		panic(err)
	}

	err = classifier.Train(trainingSet)
	if err != nil {
		panic(err)
	}

	testSet, err := csvparse.DatasetFromPath("../golearnbenchmarks/classifiers/datasets/basic_test.csv", 4, 4)
	if err != nil {
		panic(err)
	}

	totalCorrect := 0
	for i := 0; i < testSet.NumRows(); i++ {
		testRow, err := testSet.Row(i)
		if err != nil {
			panic(err)
		}

		class, err := classifier.Classify(testRow)
		if err != nil {
			panic(err)
		}

		if class.Equals(testRow.Target) {
			totalCorrect++
		}
	}
	fmt.Println(totalCorrect)
}
