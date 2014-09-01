package inmemorydataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/row"
)

type InMemoryDataset struct {
	rawDataset        []float64
	targetStart       int
	targetEnd         int
	columnTypes       []columntype.ColumnType
	allFeaturesFloats bool
	numFeatures       int
}

func NewDataset(targetStart, targetEnd int, columnTypes []columntype.ColumnType) (*InMemoryDataset, error) {
	numColumns := len(columnTypes)

	if targetOutOfBounds(targetStart, targetEnd, numColumns) {
		return nil, newTargetOutOfBoundsError(targetStart, targetEnd, numColumns)
	}

	allFeaturesFloats := true
	for i, columnType := range columnTypes {
		if (i < targetStart || i > targetEnd) && !columnType.IsFloat() {
			allFeaturesFloats = false
			break
		}
	}

	return &InMemoryDataset{[]float64{}, targetStart, targetEnd, columnTypes, allFeaturesFloats, numColumns - (targetEnd - targetStart + 1)}, nil
}

func (dataset *InMemoryDataset) NumRows() int {
	return len(dataset.rawDataset) / dataset.numColumns()
}

func (dataset *InMemoryDataset) NumFeatures() int {
	return dataset.numFeatures
}

func (dataset *InMemoryDataset) Row(i int) (*row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	numColumns := dataset.numColumns()

	target := []interface{}{}
	for j := dataset.targetStart; j <= dataset.targetEnd; j++ {
		value, err := dataset.columnTypes[j].ValueFromRaw(dataset.rawDataset[i*numColumns+j])
		if err != nil {
			return nil, err
		}

		target = append(target, value)
	}

	rawFeatureValues := dataset.rawDataset[i*numColumns : i*numColumns+dataset.targetStart]
	if dataset.targetEnd < numColumns {
		rawFeatureValues = append(rawFeatureValues, dataset.rawDataset[i*numColumns+dataset.targetEnd+1:(i+1)*numColumns]...)
	}

	return row.UnsafeNewRow(target, rawFeatureValues, dataset.AllFeaturesFloats()), nil
}

func (dataset *InMemoryDataset) AddRowFromStrings(targetStart, targetEnd int, columnTypes []columntype.ColumnType, strings []string) error {
	actualLength := len(strings)
	expectedLength := len(columnTypes)

	if actualLength != expectedLength {
		return newRowLengthMismatchError(actualLength, expectedLength)
	}

	rawValues := make([]float64, actualLength)

	for i, s := range strings {
		columnType := columnTypes[i]

		value, err := columnType.PersistRawFromString(s)
		if err != nil {
			return err
		}

		rawValues[i] = value
	}

	dataset.rawDataset = append(dataset.rawDataset, rawValues...)

	return nil
}

func (dataset *InMemoryDataset) AllFeaturesFloats() bool {
	return dataset.allFeaturesFloats
}

func (dataset *InMemoryDataset) numColumns() int {
	return len(dataset.columnTypes)
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

func newRowLengthMismatchError(actual, expected int) error {
	return errors.New(fmt.Sprintf("Row has length %d, expected %d", actual, expected))
}

func newDatasetRowIndexOutOfBoundsError(index, numRows int) error {
	return errors.New(fmt.Sprintf("Cannot access row %d in dataset with %d rows", index, numRows))
}
