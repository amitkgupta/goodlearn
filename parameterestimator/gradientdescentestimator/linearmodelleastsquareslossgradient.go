package gradientdescentestimator

import (
	"errors"
)

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
