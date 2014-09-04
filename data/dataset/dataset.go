package dataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/slice"
)

type Dataset interface {
	AllFeaturesFloats() bool
	NumFeatures() int

	AddRowFromStrings(strings []string) error
	NumRows() int
	Row(i int) (row.Row, error)
}

type inMemoryDataset struct {
	allFeaturesFloats    bool
	allTargetsFloats     bool
	featureColumnIndices []int
	targetColumnIndices  []int
	columnTypes          []columntype.ColumnType
	numFeatures          int
	numColumns           int
	rawDataset           []float64
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

	return &inMemoryDataset{
		allFeaturesFloats,
		allTargetsFloats,
		featureColumnIndices,
		targetColumnIndices,
		columnTypes,
		len(featureColumnIndices),
		len(columnTypes),
		[]float64{},
	}
}

func (dataset *inMemoryDataset) AllFeaturesFloats() bool {
	return dataset.allFeaturesFloats
}

func (dataset *inMemoryDataset) NumFeatures() int {
	return dataset.numFeatures
}

func (dataset *inMemoryDataset) AddRowFromStrings(strings []string) error {
	actualLength := len(strings)
	expectedLength := dataset.numColumns

	if actualLength != expectedLength {
		return newRowLengthMismatchError(actualLength, expectedLength)
	}

	rawValues := make([]float64, actualLength)

	for i, s := range strings {
		value, err := dataset.columnTypes[i].PersistRawFromString(s)
		if err != nil {
			return err
		}

		rawValues[i] = value
	}

	dataset.rawDataset = append(dataset.rawDataset, rawValues...)

	return nil
}

func (dataset *inMemoryDataset) NumRows() int {
	return len(dataset.rawDataset) / dataset.numColumns
}

func (dataset *inMemoryDataset) Row(i int) (row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	var rawValues []float64
	if i == numRows-1 {
		rawValues = dataset.rawDataset[i*dataset.numColumns:]
	} else {
		rawValues = dataset.rawDataset[i*dataset.numColumns : (i+1)*dataset.numColumns]
	}

	features, err := slice.SliceFromRawValues(
		dataset.allFeaturesFloats,
		dataset.featureColumnIndices,
		dataset.columnTypes,
		rawValues,
	)
	if err != nil {
		return nil, err
	}

	target, err := slice.SliceFromRawValues(
		dataset.allTargetsFloats,
		dataset.targetColumnIndices,
		dataset.columnTypes,
		rawValues,
	)
	if err != nil {
		return nil, err
	}

	return row.NewRow(features, target, dataset.numFeatures), nil
}

func newRowLengthMismatchError(actual, expected int) error {
	return errors.New(fmt.Sprintf("Row has length %d, expected %d", actual, expected))
}

func newDatasetRowIndexOutOfBoundsError(index, numRows int) error {
	return errors.New(fmt.Sprintf("Cannot access row %d in dataset with %d rows", index, numRows))
}
