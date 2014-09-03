package gradientdescent

import (
	"errors"
	"math"

	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/vectorutilities"
)

type Gradient func([]float64) ([]float64, error)

func GradientDescent(
	initialGuess []float64,
	learningRate, precision float64,
	maxIterations int,
	gradient Gradient,
) ([]float64, error) {
	// bad, check some dimensions?

	oldResult := make([]float64, len(initialGuess))
	newResult := make([]float64, len(initialGuess))
	copy(oldResult, initialGuess)

	for i := 0; i < maxIterations; i++ {
		gradientAtOldResult, err := gradient(oldResult)
		if err != nil {
			return nil, err
		}

		newResult = vectorutilities.Add(oldResult, vectorutilities.Scale(-learningRate, gradientAtOldResult))

		// bad, do this better
		if math.IsInf(newResult[0], 0) || math.IsNaN(newResult[0]) {
			return nil, errors.New("Inf or Nan")
		}

		// bad, investigate sqrt
		if (knnutilities.Euclidean(newResult, oldResult, precision)) < precision {
			return newResult, nil
		} else {
			oldResult = newResult
		}
	}

	return newResult, nil
}
