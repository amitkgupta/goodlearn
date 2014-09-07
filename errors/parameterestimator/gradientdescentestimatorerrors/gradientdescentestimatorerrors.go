package gradientdescentestimatorerrors

import (
	"fmt"
)

func NewInvalidGDPEInitializationValuesError(learningRate, precision float64, maxIterations int) InvalidGDPEInitializationValuesError {
	return InvalidGDPEInitializationValuesError{learningRate, precision, maxIterations}
}

func NewEmptyTrainingSetError() EmptyTrainingSetError {
	return EmptyTrainingSetError{}
}
func NewNonFloatFeaturesError() NonFloatFeaturesError {
	return NonFloatFeaturesError{}
}
func NewNonFloatTargetError() NonFloatTargetError {
	return NonFloatTargetError{}
}
func NewInvalidNumberOfTargetsError(numTargets int) InvalidNumberOfTargetsError {
	return InvalidNumberOfTargetsError{numTargets}
}
func NewNoFeaturesError() NoFeaturesError {
	return NoFeaturesError{}
}

func NewUntrainedEstimatorError() UntrainedEstimatorError {
	return UntrainedEstimatorError{}
}
func NewEmptyInitialParametersError() EmptyInitialParametersError {
	return EmptyInitialParametersError{}
}

type InvalidGDPEInitializationValuesError struct {
	learningRate  float64
	precision     float64
	maxIterations int
}

type EmptyTrainingSetError struct{}
type NonFloatFeaturesError struct{}
type NonFloatTargetError struct{}
type InvalidNumberOfTargetsError struct {
	numTargets int
}
type NoFeaturesError struct{}

type UntrainedEstimatorError struct{}
type EmptyInitialParametersError struct{}

func (e InvalidGDPEInitializationValuesError) Error() string {
	return fmt.Sprintf(
		"Learning rate, precision, and max iterations are %.4f, %.4f, %d; "+
			"they must all be positive",
		e.learningRate,
		e.precision,
		e.maxIterations,
	)
}

func (e EmptyTrainingSetError) Error() string {
	return "Cannot perform parameter estimation using empty dataset"
}
func (e NonFloatFeaturesError) Error() string {
	return "Cannot perform parameter estimation on dataset with non-float features"
}
func (e NonFloatTargetError) Error() string {
	return "Cannot perform parameter estimation on dataset with non-float target"
}
func (e InvalidNumberOfTargetsError) Error() string {
	return fmt.Sprintf(
		"Can only perform parameter estimation on dataset with 1 target value, got %d",
		e.numTargets,
	)
}
func (e NoFeaturesError) Error() string {
	return "Cannot perform parameter estimation on dataset with no feature values"
}

func (e UntrainedEstimatorError) Error() string {
	return "Cannot perform estimation with an untrained estimator"
}
func (e EmptyInitialParametersError) Error() string {
	return "Initial parameters must not be empty"
}
