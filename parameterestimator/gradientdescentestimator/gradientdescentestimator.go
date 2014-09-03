package gradientdescentestimator

import (
	"errors"

	"github.com/amitkgupta/goodlearn/optimizer/gradientdescent"
	"github.com/amitkgupta/goodlearn/vectorutilities"
)

type ParameterizedLossGradient func([]float64, []float64, float64) ([]float64, error)

func GradientDescentParameterEstimation(
	initialParameters, observedX, observedY []float64,
	learningRate, precision float64,
	maxIterations int,
	plgf ParameterizedLossGradient,
) ([]float64, error) {
	// bad, handle problems with dimensions, noting that x dimension and parameters need not match

	xDimension := len(observedX) / len(observedY)

	gradient := func(x []float64) ([]float64, error) {
		sumLossGradient := make([]float64, len(initialParameters))

		for i, y := range observedY {
			lossGradient, err := plgf(x, observedX[i*xDimension:(i+1)*xDimension], y)
			if err != nil {
				return nil, err
			}
			sumLossGradient = vectorutilities.Add(sumLossGradient, lossGradient)
		}

		return sumLossGradient, nil
	}

	return gradientdescent.GradientDescent(initialParameters, learningRate, precision, maxIterations, gradient)
}

func LinearModelLeastSquaresLossGradient(parameters, observedX []float64, observedY float64) ([]float64, error) {
	if len(parameters) != len(observedX)+1 {
		return nil, errors.New("need exactly one more parameter than observed Xs for the constant term")
	}

	z := parameters[len(parameters)-1]
	for i, x := range observedX {
		z = z + parameters[i]*x
	}
	z = 2 * (z - observedY)

	result := make([]float64, len(parameters))
	result[len(parameters)-1] = z
	for i := range parameters[:len(parameters)-1] {
		result[i] = observedX[i] * z
	}

	return result, nil
}
