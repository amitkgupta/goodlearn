package dataset

import (
	"errors"
	"fmt"

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

func NewDataset(targetStartInclusive, targetEndExclusive int, columnTypes []columntype.ColumnType) (Dataset, error) {
	numColumns := len(columnTypes)

	if targetOutOfBounds(targetStartInclusive, targetEndExclusive, numColumns) {
		return nil, newTargetOutOfBoundsError(targetStartInclusive, targetEndExclusive, numColumns)
	}

	featureColumnIndices := featureColumnIndices(targetStartInclusive, targetEndExclusive, numColumns)
	targetColumnIndices := targetColumnIndices(targetStartInclusive, targetEndExclusive, numColumns)

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
	), nil
}

func featureColumnIndices(targetStartInclusive, targetEndExclusive, numColumns int) []int {
	result := []int{}
	for i := 0; i < numColumns; i++ {
		if i < targetStartInclusive || i >= targetEndExclusive {
			result = append(result, i)
		}
	}
	return result
}

func targetColumnIndices(targetStartInclusive, targetEndExclusive, numColumns int) []int {
	result := []int{}
	for i := 0; i < numColumns; i++ {
		if i >= targetStartInclusive && i < targetEndExclusive {
			result = append(result, i)
		}
	}
	return result
}

func targetOutOfBounds(targetStartInclusive, targetEndExclusive, numColumns int) bool {
	return targetStartInclusive < 0 ||
		targetEndExclusive > numColumns ||
		targetStartInclusive >= targetEndExclusive ||
		(targetEndExclusive-targetStartInclusive) >= (numColumns)
}

func newTargetOutOfBoundsError(targetStartInclusive, targetEndExclusive, numColumns int) error {
	return errors.New(fmt.Sprintf(
		"Dataset must have valid target bounds, and at least one non-target column; "+
			"cannot have %d total columns, target start column %d and target end column %d",
		numColumns,
		targetStartInclusive,
		targetEndExclusive,
	))
}
