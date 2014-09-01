package classifier

import (
	"github.com/amitkgupta/goodlearn/dataset/dataset"
	"github.com/amitkgupta/goodlearn/dataset/row"
	"github.com/amitkgupta/goodlearn/dataset/target"
)

type Classifier interface {
	Train(*dataset.Dataset) error
	Classify(*row.Row) (target.Target, error)
}
