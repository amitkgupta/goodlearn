package inmemorydataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/row"
)

type inMemoryDataset struct {
	allFeaturesFloats    bool
	allTargetsFloats     bool
	featureColumnIndices []int
	targetColumnIndices  []int
	columnTypes          []columntype.ColumnType
	rawDataset           []float64
}

func NewDataset(allFeaturesFloats, allTargetsFloats bool, featureColumnIndices, targetColumnIndices []int, columnTypes []columntype.ColumnType) *inMemoryDataset {
	return &inMemoryDataset{
		allFeaturesFloats,
		allTargetsFloats,
		featureColumnIndices,
		targetColumnIndices,
		columnTypes,
		[]float64{},
	}
}

func (dataset *inMemoryDataset) AllFeaturesFloats() bool {
	return dataset.allFeaturesFloats
}

func (dataset *inMemoryDataset) NumFeatures() int {
	return len(dataset.featureColumnIndices)
}

func (dataset *inMemoryDataset) AddRowFromStrings(strings []string) error {
	actualLength := len(strings)
	expectedLength := dataset.numColumns()

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
	return len(dataset.rawDataset) / dataset.numColumns()
}

func (dataset *inMemoryDataset) Row(i int) (row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	target := make([]interface{}, len(dataset.targetColumnIndices))
	for idx, j := range dataset.targetColumnIndices {
		target[idx] = dataset.valueAt(i, j)
	}

	rawFeatureValues := make([]float64, dataset.NumFeatures())
	for idx, j := range dataset.featureColumnIndices {
		rawFeatureValues[idx] = dataset.rawValueAt(i, j)
	}

	return row.NewRow(dataset.AllFeaturesFloats(), target, rawFeatureValues), nil
}

func (dataset *inMemoryDataset) numColumns() int {
	return len(dataset.columnTypes)
}

func (dataset *inMemoryDataset) valueAt(rowIndex, columnIndex int) interface{} {
	columnType := dataset.columnTypes[columnIndex]
	rawValue := dataset.rawValueAt(rowIndex, columnIndex)

	if floatColumnType, ok := columnType.(columntype.FloatColumnType); ok {
		value, _ := floatColumnType.ValueFromRaw(rawValue)
		return value
	} else if stringColumnType, ok := columnType.(columntype.StringColumnType); ok {
		value, _ := stringColumnType.ValueFromRaw(rawValue)
		return value
	} else {
		// bad
		panic("Unknown column type")
		return nil
	}
}

func (dataset *inMemoryDataset) rawValueAt(rowIndex, columnIndex int) float64 {
	return dataset.rawDataset[rowIndex*dataset.numColumns()+columnIndex]
}

func newUnknownColumnTypeError(columnIndex int) error {
	return errors.New(fmt.Sprintf("Column %d has unknown type", columnIndex))
}

func newRowLengthMismatchError(actual, expected int) error {
	return errors.New(fmt.Sprintf("Row has length %d, expected %d", actual, expected))
}

func newDatasetRowIndexOutOfBoundsError(index, numRows int) error {
	return errors.New(fmt.Sprintf("Cannot access row %d in dataset with %d rows", index, numRows))
}
