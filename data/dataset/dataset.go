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

	AddRowFromStrings(targetStart, targetEnd int, columnTypes []columntype.ColumnType, strings []string) error
	NumRows() int
	Row(i int) (row.Row, error)
}

func NewDataset(targetStart, targetEnd int, columnTypes []columntype.ColumnType) (Dataset, error) {
	numColumns := len(columnTypes)

	if targetOutOfBounds(targetStart, targetEnd, numColumns) {
		return nil, newTargetOutOfBoundsError(targetStart, targetEnd, numColumns)
	}

	allFeaturesFloats := true
	for i, columnType := range columnTypes {
		_, ok := columnType.(columntype.FloatColumnType)
		if !ok && (i < targetStart || i > targetEnd) {
			allFeaturesFloats = false
			break
		}
	}

	return inmemorydataset.NewDataset(
		allFeaturesFloats,
		targetStart,
		targetEnd,
		columnTypes,
	), nil
}

func targetOutOfBounds(targetStart, targetEnd, numColumns int) bool {
	return targetStart < 0 ||
		targetEnd >= numColumns ||
		targetStart > targetEnd ||
		(targetEnd-targetStart) >= (numColumns-1)
}

func newTargetOutOfBoundsError(targetStart, targetEnd, numColumns int) error {
	return errors.New(fmt.Sprintf(
		"Dataset must have valid target bounds, and at least one non-target column; "+
			"cannot have %d total columns, target start column %d and target end column %d",
		numColumns,
		targetStart,
		targetEnd,
	))
}
