package knn

import (
	"fmt"

	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"
)

func NewKNNClassifier(k int) (*kNNClassifier, error) {
	if k < 1 {
		return nil, newInvalidNumberOfNeighboursError(k)
	}

	return &kNNClassifier{k: k}, nil
}

type kNNClassifier struct {
	k            int
	trainingData dataset.FloatFeatureDataset
}

func (classifier *kNNClassifier) Train(trainingData dataset.Dataset) error {
	floatFeatureTrainingData, ok := trainingData.(dataset.FloatFeatureDataset)
	if !ok {
		return newNonFloatFeaturesTrainingSetError()
	}

	if floatFeatureTrainingData.NumRows() == 0 {
		return newEmptyTrainingDatasetError()
	}

	classifier.trainingData = floatFeatureTrainingData
	return nil
}

func (classifier *kNNClassifier) Classify(testRow row.Row) (target.Target, error) {
	floatFeatureTestRow, ok := testRow.(row.FloatFeatureRow)
	if !ok {
		return nil, newNonFloatFeaturesTestRowError()
	}

	trainingData := classifier.trainingData
	if trainingData == nil {
		return nil, newUntrainedClassifierError()
	}

	numTestRowFeatures := floatFeatureTestRow.NumFeatures()
	numTrainingDataFeatures := trainingData.NumFeatures()
	if numTestRowFeatures != numTrainingDataFeatures {
		return nil, newRowLengthMismatchError(numTestRowFeatures, numTrainingDataFeatures)
	}

	nearestNeighbours, err := knnutilities.NewKNNTargetCollection(classifier.k)
	if err != nil {
		return nil, err
	}

	testRowFeatureValues := floatFeatureTestRow.FloatFeatureValues()

	for i := 0; i < trainingData.NumRows(); i++ {
		trainingRow, err := trainingData.FloatFeatureRow(i)
		if err != nil {
			return nil, err
		}
		trainingRowFeatureValues := trainingRow.FloatFeatureValues()

		distance := knnutilities.Euclidean(testRowFeatureValues, trainingRowFeatureValues, nearestNeighbours.MaxDistance())
		if distance < nearestNeighbours.MaxDistance() {
			nearestNeighbours.Insert(trainingRow.Target(), distance)
		}
	}

	return nearestNeighbours.Vote()
}

func newInvalidNumberOfNeighboursError(k int) InvalidNumberOfNeighboursError {
	return InvalidNumberOfNeighboursError{k}
}

func newEmptyTrainingDatasetError() EmptyTrainingDatasetError {
	return EmptyTrainingDatasetError{}
}
func newNonFloatFeaturesTrainingSetError() NonFloatFeaturesTrainingSetError {
	return NonFloatFeaturesTrainingSetError{}
}

func newUntrainedClassifierError() UntrainedClassifierError {
	return UntrainedClassifierError{}
}
func newNonFloatFeaturesTestRowError() NonFloatFeaturesTestRowError {
	return NonFloatFeaturesTestRowError{}
}
func newRowLengthMismatchError(numTestRowFeatures, numTrainingSetFeatures int) RowLengthMismatchError {
	return RowLengthMismatchError{numTestRowFeatures, numTrainingSetFeatures}
}

type InvalidNumberOfNeighboursError struct {
	k int
}

type EmptyTrainingDatasetError struct{}
type NonFloatFeaturesTrainingSetError struct {
	kNNClassifier
}

type UntrainedClassifierError struct{}
type NonFloatFeaturesTestRowError struct{}
type RowLengthMismatchError struct {
	numTestRowFeatures     int
	numTrainingSetFeatures int
}

func (e InvalidNumberOfNeighboursError) Error() string {
	return fmt.Sprintf("invalid number of neighbours %d", e.k)
}

func (e EmptyTrainingDatasetError) Error() string {
	return "cannot train on an empty dataset"
}
func (e NonFloatFeaturesTrainingSetError) Error() string {
	return "cannot train on dataset with some non-float features"
}

func (e NonFloatFeaturesTestRowError) Error() string {
	return "cannot classify a row with some non-float features"
}
func (e UntrainedClassifierError) Error() string {
	return "cannot classify before training"
}
func (e RowLengthMismatchError) Error() string {
	return fmt.Sprintf("Test row has %d features, training set has %d", e.numTestRowFeatures, e.numTrainingSetFeatures)
}
