package dataset

import (
	"github.com/amitkgupta/goodlearn/data/row"
)

type Dataset interface {
	NumFeatures() int
	NumRows() int
	AddRowFromStrings(targetStart, targetEnd int, strings []string) error
}

type MixedFeatureDataset interface {
	Dataset
	Row(i int) (row.Row, error)
}

type FloatFeatureDataset interface {
	Dataset
	FloatFeatureRow(i int) (row.FloatFeatureRow, error)
}
