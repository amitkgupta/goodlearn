package gradientdescentestimator

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/optimizer/gradientdescent"
	"github.com/amitkgupta/goodlearn/vectorutilities"
)

type ParameterizedLossGradient func([]float64, []float64, float64) ([]float64, error)

type gradientDescentParameterEstimator struct {
	learningRate  float64
	precision     float64
	maxIterations int
	plgf          ParameterizedLossGradient
	trainingSet   dataset.Dataset
}

func NewGradientDescentParameterEstimator(
	learningRate, precision float64,
	maxIterations int,
	plgf ParameterizedLossGradient,
) (*gradientDescentParameterEstimator, error) {
	if learningRate <= 0 || precision <= 0 || maxIterations <= 0 {
		return nil, errors.New(fmt.Sprintf(
			"Learning rate, precision, and max iterations are %.4f, %.4f, %d; "+
				"they must all be positive",
			learningRate,
			precision,
			maxIterations,
		))
	}

	return &gradientDescentParameterEstimator{
		learningRate:  learningRate,
		precision:     precision,
		maxIterations: maxIterations,
		plgf:          plgf,
	}, nil
}

func (gdpe *gradientDescentParameterEstimator) Train(ds dataset.Dataset) error {
	if ds.NumRows() == 0 {
		return errors.New("Cannot perform estimation using empty dataset")
	}

	if !ds.AllFeaturesFloats() {
		return errors.New("Cannot perform parameter estimation on dataset with non-float features")
	}

	if !ds.AllTargetsFloats() {
		return errors.New("Cannot perform parameter estimation on dataset with non-float target")
	}

	if ds.NumTargets() != 1 {
		return errors.New("Cannot perform parameter estimation on dataset with 1 target value")
	}

	if ds.NumFeatures() == 0 {
		return errors.New("Cannot perform parameter estimation on dataset with no feature values")
	}

	gdpe.trainingSet = ds
	return nil
}

func (gdpe *gradientDescentParameterEstimator) Estimate(initialParameters []float64) ([]float64, error) {
	if gdpe.trainingSet == nil {
		return nil, errors.New("Cannot perform estimation with an untrained estimator")
	}

	if len(initialParameters) == 0 {
		return nil, errors.New("Initial parameters must not be empty")
	}

	gradient := func(guess []float64) ([]float64, error) {
		sumLossGradient := make([]float64, len(initialParameters))

		for i := 0; i < gdpe.trainingSet.NumRows(); i++ {
			row, _ := gdpe.trainingSet.Row(i)
			features, _ := row.Features().(slice.FloatSlice)
			target, _ := row.Target().(slice.FloatSlice)
			x := features.Values()
			y := target.Values()[0]

			lossGradient, err := gdpe.plgf(guess, x, y)
			if err != nil {
				return nil, err
			}
			sumLossGradient = vectorutilities.Add(sumLossGradient, lossGradient)
		}

		return sumLossGradient, nil
	}

	return gradientdescent.GradientDescent(initialParameters, gdpe.learningRate, gdpe.precision, gdpe.maxIterations, gradient)
}
