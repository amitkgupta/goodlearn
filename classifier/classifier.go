package classifier

import (
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"
)

type Classifier interface {
	Train(dataset.Dataset) error
	Classify(row.Row) (target.Target, error)
}
