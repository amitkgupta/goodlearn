package gradientdescentestimator

import (
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/slice"
	"github.com/amitkgupta/goodlearn/optimizer/gradientdescent"
	gdeUtilities "github.com/amitkgupta/goodlearn/parameterestimator/gradientdescentestimator/gradientdescentestimatorutilities"
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
		return nil, gdeUtilities.NewInvalidGDPEInitializationValuesError(learningRate, precision, maxIterations)
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
		return gdeUtilities.NewEmptyTrainingSetError()
	}

	if !ds.AllFeaturesFloats() {
		return gdeUtilities.NewNonFloatFeaturesError()
	}

	if !ds.AllTargetsFloats() {
		return gdeUtilities.NewNonFloatTargetError()
	}

	if ds.NumTargets() != 1 {
		return gdeUtilities.NewInvalidNumberOfTargetsError(ds.NumTargets())
	}

	if ds.NumFeatures() == 0 {
		return gdeUtilities.NewNoFeaturesError()
	}

	gdpe.trainingSet = ds
	return nil
}

func (gdpe *gradientDescentParameterEstimator) Estimate(initialParameters []float64) ([]float64, error) {
	if gdpe.trainingSet == nil {
		return nil, gdeUtilities.NewUntrainedEstimatorError()
	}

	if len(initialParameters) == 0 {
		return nil, gdeUtilities.NewEmptyInitialParametersError()
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
