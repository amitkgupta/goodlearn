package dataset

import (
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset/inmemorydataset"
	"github.com/amitkgupta/goodlearn/data/row"
)

type Dataset interface {
	AllFeaturesFloats() bool
	NumFeatures() int

	AddRowFromStrings(strings []string) error
	NumRows() int
	Row(i int) (row.Row, error)
}

func NewDataset(featureColumnIndices, targetColumnIndices []int, columnTypes []columntype.ColumnType) Dataset {
	allFeaturesFloats := true
	for _, i := range featureColumnIndices {
		if _, ok := columnTypes[i].(columntype.FloatColumnType); !ok {
			allFeaturesFloats = false
			break
		}
	}

	allTargetsFloats := true
	for _, i := range targetColumnIndices {
		if _, ok := columnTypes[i].(columntype.FloatColumnType); !ok {
			allTargetsFloats = false
			break
		}
	}

	return inmemorydataset.NewDataset(
		allFeaturesFloats,
		allTargetsFloats,
		featureColumnIndices,
		targetColumnIndices,
		columnTypes,
	)
}
