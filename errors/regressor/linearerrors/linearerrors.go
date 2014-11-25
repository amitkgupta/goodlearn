package linearerrors

import (
	"fmt"
)

func NewNonFloatFeaturesError() NonFloatFeaturesTrainingSetError {
	return NonFloatFeaturesTrainingSetError{}
}
func NewNonFloatTargetsError() NonFloatTargetsTrainingSetError {
	return NonFloatTargetsTrainingSetError{}
}
func NewInvalidNumberOfTargetsError(numTargets int) InvalidNumberOfTargetsError {
	return InvalidNumberOfTargetsError{numTargets}
}
func NewNoFeaturesError() NoFeaturesError {
	return NoFeaturesError{}
}
func NewEstimatorConstructionError(err error) EstimatorConstructionError {
	return EstimatorConstructionError{err}
}
func NewEstimatorTrainingError(err error) EstimatorTrainingError {
	return EstimatorTrainingError{err}
}
func NewEstimatorEstimationError(err error) EstimatorEstimationError {
	return EstimatorEstimationError{err}
}

func NewUntrainedRegressorError() UntrainedRegressorError {
	return UntrainedRegressorError{}
}
func NewRowLengthMismatchError(numTestRowFeatures, numTrainingSetFeatures int) RowLengthMismatchError {
	return RowLengthMismatchError{numTestRowFeatures, numTrainingSetFeatures}
}
func NewNonFloatFeaturesTestRowError() NonFloatFeaturesTestRowError {
	return NonFloatFeaturesTestRowError{}
}

type NonFloatFeaturesTrainingSetError struct{}
type NonFloatTargetsTrainingSetError struct{}
type InvalidNumberOfTargetsError struct {
	numTargets int
}
type NoFeaturesError struct{}
type EstimatorConstructionError struct {
	err error
}
type EstimatorTrainingError struct {
	err error
}
type EstimatorEstimationError struct {
	err error
}
type UntrainedRegressorError struct{}
type RowLengthMismatchError struct {
	numTestRowFeatures     int
	numTrainingSetFeatures int
}
type NonFloatFeaturesTestRowError struct{}

func (e NonFloatFeaturesTrainingSetError) Error() string {
	return "cannot train on dataset with some non-float features"
}
func (e NonFloatTargetsTrainingSetError) Error() string {
	return "cannot train on dataset with some non-float targets"
}
func (e InvalidNumberOfTargetsError) Error() string {
	return fmt.Sprintf("cannot train regressor on dataset with %d targets, must have exactly 1", e.numTargets)
}
func (e NoFeaturesError) Error() string {
	return "cannot train regressor on dataset with no features"
}
func (e EstimatorConstructionError) Error() string {
	return fmt.Sprintf("could not construct estimator: %s", e.err.Error())
}
func (e EstimatorTrainingError) Error() string {
	return fmt.Sprintf("could not train estimator: %s", e.err.Error())
}
func (e EstimatorEstimationError) Error() string {
	return fmt.Sprintf("could not estimate coefficients: %s", e.err.Error())
}

func (e UntrainedRegressorError) Error() string {
	return "cannot predict before training"
}
func (e RowLengthMismatchError) Error() string {
	return fmt.Sprintf("Test row has %d features, training set has %d", e.numTestRowFeatures, e.numTrainingSetFeatures)
}
func (e NonFloatFeaturesTestRowError) Error() string {
	return "cannot predict row with some non-float features"
}
