package dataset

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/row"
)

type Dataset interface {
	AllFeaturesFloats() bool
	NumFeatures() int
	NumRows() int
	Row(i int) (*row.Row, error)
	AddRowFromStrings(targetStart, targetEnd int, columnTypes []columntype.ColumnType, strings []string) error
}
