package parameterestimator

import (
	"github.com/amitkgupta/goodlearn/data/dataset"
)

type ParameterEstimator interface {
	Train(dataset.Dataset) error
	Estimate([]float64) ([]float64, error)
}
