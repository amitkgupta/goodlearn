package knn

import (
	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"
)

func NewKNNClassifier(k int) (*kNNClassifier, error) {
	if k < 1 {
		return nil, knnutilities.NewInvalidNumberOfNeighboursError(k)
	}

	return &kNNClassifier{k: k}, nil
}

type kNNClassifier struct {
	k            int
	trainingData dataset.Dataset
}

func (classifier *kNNClassifier) Train(trainingData dataset.Dataset) error {
	if !trainingData.AllFeaturesFloats() {
		return knnutilities.NewNonFloatFeaturesTrainingSetError()
	}

	if trainingData.NumRows() == 0 {
		return knnutilities.NewEmptyTrainingDatasetError()
	}

	classifier.trainingData = trainingData
	return nil
}

func (classifier *kNNClassifier) Classify(testRow row.Row) (target.Target, error) {
	trainingData := classifier.trainingData
	if trainingData == nil {
		return nil, knnutilities.NewUntrainedClassifierError()
	}

	numTestRowFeatures := testRow.NumFeatures()
	numTrainingDataFeatures := trainingData.NumFeatures()
	if numTestRowFeatures != numTrainingDataFeatures {
		return nil, knnutilities.NewRowLengthMismatchError(numTestRowFeatures, numTrainingDataFeatures)
	}

	floatFeatureTestRow, ok := testRow.(row.FloatFeatureRow)
	if !ok {
		return nil, knnutilities.NewNonFloatFeaturesTestRowError()
	}
	testRowFeatureValues := floatFeatureTestRow.Features()

	nearestNeighbours := knnutilities.NewKNNTargetCollection(classifier.k)

	for i := 0; i < trainingData.NumRows(); i++ {
		trainingRow, _ := trainingData.Row(i)
		floatFeatureTrainingRow, _ := trainingRow.(row.FloatFeatureRow)
		trainingRowFeatureValues := floatFeatureTrainingRow.Features()

		distance := knnutilities.Euclidean(testRowFeatureValues, trainingRowFeatureValues, nearestNeighbours.MaxDistance())
		if distance < nearestNeighbours.MaxDistance() {
			nearestNeighbours.Insert(trainingRow.Target(), distance)
		}
	}

	return nearestNeighbours.Vote(), nil
}
