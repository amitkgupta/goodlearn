package knn

import (
	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/slice"
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

func (classifier *kNNClassifier) Classify(testRow row.Row) (slice.Slice, error) {
	trainingData := classifier.trainingData
	if trainingData == nil {
		return nil, knnutilities.NewUntrainedClassifierError()
	}

	numTestRowFeatures := testRow.NumFeatures()
	numTrainingDataFeatures := trainingData.NumFeatures()
	if numTestRowFeatures != numTrainingDataFeatures {
		return nil, knnutilities.NewRowLengthMismatchError(numTestRowFeatures, numTrainingDataFeatures)
	}

	testFeatures, ok := testRow.Features().(slice.FloatSlice)
	if !ok {
		return nil, knnutilities.NewNonFloatFeaturesTestRowError()
	}
	testFeatureValues := testFeatures.Values()

	nearestNeighbours := knnutilities.NewKNNTargetCollection(classifier.k)

	for i := 0; i < trainingData.NumRows(); i++ {
		trainingRow, _ := trainingData.Row(i)
		trainingFeatures, _ := trainingRow.Features().(slice.FloatSlice)
		trainingFeatureValues := trainingFeatures.Values()

		distance := knnutilities.Euclidean(testFeatureValues, trainingFeatureValues, nearestNeighbours.MaxDistance())
		if distance < nearestNeighbours.MaxDistance() {
			nearestNeighbours.Insert(trainingRow.Target(), distance)
		}
	}

	return nearestNeighbours.Vote(), nil
}
