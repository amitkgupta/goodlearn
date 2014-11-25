package regressor

import (
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
)

type Regressor interface {
	Train(dataset.Dataset) error
	Predict(row.Row) (float64, error)
}
