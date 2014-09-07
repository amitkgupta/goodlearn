package knnerrors

import (
	"fmt"
)

func NewInvalidNumberOfNeighboursError(k int) InvalidNumberOfNeighboursError {
	return InvalidNumberOfNeighboursError{k}
}

func NewEmptyTrainingDatasetError() EmptyTrainingDatasetError {
	return EmptyTrainingDatasetError{}
}
func NewNonFloatFeaturesTrainingSetError() NonFloatFeaturesTrainingSetError {
	return NonFloatFeaturesTrainingSetError{}
}

func NewUntrainedClassifierError() UntrainedClassifierError {
	return UntrainedClassifierError{}
}
func NewRowLengthMismatchError(numTestRowFeatures, numTrainingSetFeatures int) RowLengthMismatchError {
	return RowLengthMismatchError{numTestRowFeatures, numTrainingSetFeatures}
}
func NewNonFloatFeaturesTestRowError() NonFloatFeaturesTestRowError {
	return NonFloatFeaturesTestRowError{}
}

type InvalidNumberOfNeighboursError struct {
	k int
}

type EmptyTrainingDatasetError struct{}
type NonFloatFeaturesTrainingSetError struct{}

type UntrainedClassifierError struct{}
type RowLengthMismatchError struct {
	numTestRowFeatures     int
	numTrainingSetFeatures int
}
type NonFloatFeaturesTestRowError struct{}

func (e InvalidNumberOfNeighboursError) Error() string {
	return fmt.Sprintf("invalid number of neighbours %d", e.k)
}

func (e EmptyTrainingDatasetError) Error() string {
	return "cannot train on an empty dataset"
}
func (e NonFloatFeaturesTrainingSetError) Error() string {
	return "cannot train on dataset with some non-float features"
}

func (e UntrainedClassifierError) Error() string {
	return "cannot classify before training"
}
func (e RowLengthMismatchError) Error() string {
	return fmt.Sprintf("Test row has %d features, training set has %d", e.numTestRowFeatures, e.numTrainingSetFeatures)
}
func (e NonFloatFeaturesTestRowError) Error() string {
	return "cannot classify row with some non-float features"
}
