package linear

import (
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/errors/regressor/linearerrors"
	"github.com/amitkgupta/goodlearn/parameterestimator/gradientdescentestimator"
)

func NewLinearRegressor() *linearRegressor {
	return &linearRegressor{}
}

type linearRegressor struct {
	coefficients []float64
}

const (
	defaultLearningRate  = 0.004
	defaultPrecision     = 1e-8
	defaultMaxIterations = 1e8
)

func (regressor *linearRegressor) Train(trainingData dataset.Dataset) error {
	if !trainingData.AllFeaturesFloats() {
		return linearerrors.NewNonFloatFeaturesError()
	}

	if !trainingData.AllTargetsFloats() {
		return linearerrors.NewNonFloatTargetsError()
	}

	if trainingData.NumTargets() != 1 {
		return linearerrors.NewInvalidNumberOfTargetsError(trainingData.NumTargets())
	}

	if trainingData.NumFeatures() == 0 {
		return linearerrors.NewNoFeaturesError()
	}

	estimator, err := gradientdescentestimator.NewGradientDescentParameterEstimator(
		defaultLearningRate,
		defaultPrecision,
		defaultMaxIterations,
		gradientdescentestimator.LinearModelLeastSquaresLossGradient,
	)
	if err != nil {
		return linearerrors.NewEstimatorConstructionError(err)
	}

	err = estimator.Train(trainingData)
	if err != nil {
		return linearerrors.NewEstimatorTrainingError(err)
	}

	coefficients, err := estimator.Estimate(defaultInitialCoefficientEstimate(trainingData.NumFeatures()))
	if err != nil {
		return linearerrors.NewEstimatorEstimationError(err)
	}

	regressor.coefficients = coefficients
	return nil
}

func (regressor *linearRegressor) Predict(testRow row.Row) (float64, error) {
	coefficients := regressor.coefficients
	if coefficients == nil {
		return 0, linearerrors.NewUntrainedRegressorError()
	}

	numTestRowFeatures := testRow.NumFeatures()
	numCoefficients := len(coefficients)
	if numCoefficients != numTestRowFeatures+1 {
		return 0, linearerrors.NewRowLengthMismatchError(numTestRowFeatures, numCoefficients)
	}

	testFeatures, ok := testRow.Features().(slice.FloatSlice)
	if !ok {
		return 0, linearerrors.NewNonFloatFeaturesTestRowError()
	}
	testFeatureValues := testFeatures.Values()

	result := coefficients[numCoefficients-1]
	for i, c := range coefficients[:numCoefficients-1] {
		result = result + c*testFeatureValues[i]
	}

	return result, nil
}

func defaultInitialCoefficientEstimate(numFeatures int) []float64 {
	return make([]float64, numFeatures+1)
}
