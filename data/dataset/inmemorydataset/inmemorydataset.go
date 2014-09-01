package inmemorydataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/row"
)

type inMemoryDataset struct {
	allFeaturesFloats bool
	numFeatures       int

	targetStart int
	targetEnd   int
	columnTypes []columntype.ColumnType
	rawDataset  []float64
}

func NewDataset(allFeaturesFloats bool, targetStart, targetEnd int, columnTypes []columntype.ColumnType) *inMemoryDataset {
	numColumns := len(columnTypes)
	numFeatures := numColumns - (targetEnd - targetStart + 1)
	return &inMemoryDataset{
		allFeaturesFloats,
		numFeatures,
		targetStart,
		targetEnd,
		columnTypes,
		[]float64{},
	}
}

func (dataset *inMemoryDataset) AllFeaturesFloats() bool {
	return dataset.allFeaturesFloats
}

func (dataset *inMemoryDataset) NumFeatures() int {
	return dataset.numFeatures
}

func (dataset *inMemoryDataset) AddRowFromStrings(targetStart, targetEnd int, columnTypes []columntype.ColumnType, strings []string) error {
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

func (dataset *inMemoryDataset) NumRows() int {
	return len(dataset.rawDataset) / dataset.numColumns()
}

func (dataset *inMemoryDataset) Row(i int) (row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	numColumns := dataset.numColumns()

	target := []interface{}{}
	for j := dataset.targetStart; j <= dataset.targetEnd; j++ {
		value, err := dataset.valueAt(i, j)
		if err != nil {
			return nil, err
		}
		target = append(target, value)
	}

	rawFeatureValues := dataset.rawDataset[i*numColumns : i*numColumns+dataset.targetStart]
	if dataset.targetEnd < numColumns {
		rawFeatureValues = append(rawFeatureValues, dataset.rawDataset[i*numColumns+dataset.targetEnd+1:(i+1)*numColumns]...)
	}

	return row.NewRow(dataset.AllFeaturesFloats(), target, rawFeatureValues), nil
}

func (dataset *inMemoryDataset) numColumns() int {
	return len(dataset.columnTypes)
}

func (dataset *inMemoryDataset) valueAt(rowIndex, columnIndex int) (interface{}, error) {
	columnType := dataset.columnTypes[columnIndex]
	rawValue := dataset.rawDataset[rowIndex*dataset.numColumns()+columnIndex]

	var value interface{}
	var err error

	if floatColumnType, ok := columnType.(columntype.FloatColumnType); ok {
		value, err = floatColumnType.ValueFromRaw(rawValue)
	} else if stringColumnType, ok := columnType.(columntype.StringColumnType); ok {
		value, err = stringColumnType.ValueFromRaw(rawValue)
	} else {
		value, err = nil, newUnknownColumnTypeError(columnIndex)
	}

	return value, err
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
