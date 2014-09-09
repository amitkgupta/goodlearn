package gradientdescent

import (
	"errors"

	"github.com/amitkgupta/goodlearn/classifier/knn/knnutilities"
	"github.com/amitkgupta/goodlearn/vectorutilities"
)

func GradientDescent(
	initialGuess []float64,
	learningRate, precision float64,
	maxIterations int,
	gradient func([]float64) ([]float64, error),
) ([]float64, error) {
	if len(initialGuess) == 0 {
		return nil, errors.New("initialGuess cannot be empty")
	}

	oldResult := make([]float64, len(initialGuess))
	newResult := make([]float64, len(initialGuess))
	copy(oldResult, initialGuess)

	for i := 0; i < maxIterations; i++ {
		gradientAtOldResult, err := gradient(oldResult)
		if err != nil {
			return nil, err
		}

		newResult = vectorutilities.Add(oldResult, vectorutilities.Scale(-learningRate, gradientAtOldResult))

		if (knnutilities.Euclidean(newResult, oldResult, precision)) < precision*precision {
			return newResult, nil
		} else {
			oldResult = newResult
		}
	}

	return newResult, nil
}
