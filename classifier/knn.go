package classifier

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/classifier/knnutilities"
	"github.com/amitkgupta/goodlearn/dataset/dataset"
	"github.com/amitkgupta/goodlearn/dataset/row"
	"github.com/amitkgupta/goodlearn/dataset/target"
)

func NewKNNClassifier(k int) (*kNNClassifier, error) {
	if k < 1 {
		return nil, newInvalidNumberOfNeighboursError(k)
	}

	return &kNNClassifier{k: k}, nil
}

type kNNClassifier struct {
	k            int
	trainingData *dataset.Dataset
}

// bad, check compaitiblity
func (classifier *kNNClassifier) Train(trainingData *dataset.Dataset) error {
	classifier.trainingData = trainingData
	return nil
}

func (classifier *kNNClassifier) Classify(testRow *row.Row) (target.Target, error) {
	trainingData := classifier.trainingData
	if trainingData == nil {
		return nil, newUntrainedClassifierError()
	}

	// bad, check compatibility

	nearestNeighbours := knnutilities.NewSortedTargetCollection(classifier.k)
	testRowFeatureValues := testRow.UnsafeFloatFeatureValues()

	for i := 0; i < trainingData.NumRows(); i++ {
		trainingRow, err := trainingData.Row(i)
		if err != nil {
			return nil, err
		}
		trainingRowFeatureValues := trainingRow.UnsafeFloatFeatureValues()

		distance := knnutilities.Euclidean(testRowFeatureValues, trainingRowFeatureValues, nearestNeighbours.MaxDistance())
		if distance < nearestNeighbours.MaxDistance() {
			nearestNeighbours.Insert(trainingRow.Target, distance)
		}
	}

	return nearestNeighbours.Vote()
}

func newUntrainedClassifierError() error {
	return errors.New("cannot classify before training")
}

func newInvalidNumberOfNeighboursError(k int) error {
	return errors.New(fmt.Sprintf("invalid number of neighbours %d", k))
}

func newInvalidFloatFeatureDatasetError() error {
	return errors.New("dataset invalid, has feature columns which aren't floats")
}
