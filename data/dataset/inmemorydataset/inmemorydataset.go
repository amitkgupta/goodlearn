package inmemorydataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/row/inmemoryrow"
	"github.com/amitkgupta/goodlearn/data/target"
)

type mixedFeatureTypeDataset struct {
	rawDataset  []float64
	targetStart int
	targetEnd   int
	columnTypes []columntype.ColumnType
	numFeatures int
}

type floatFeatureTypeDataset struct {
	*mixedFeatureTypeDataset
}

func NewDataset(targetStart, targetEnd int, columnTypes []columntype.ColumnType) (dataset.Dataset, error) {
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

	mixedFeatureTypeDataset := &mixedFeatureTypeDataset{
		[]float64{},
		targetStart,
		targetEnd,
		columnTypes,
		numColumns - (targetEnd - targetStart + 1),
	}

	if allFeaturesFloats {
		return &floatFeatureTypeDataset{mixedFeatureTypeDataset}, nil
	} else {
		return mixedFeatureTypeDataset, nil
	}
}

func (dataset *mixedFeatureTypeDataset) NumRows() int {
	return len(dataset.rawDataset) / dataset.numColumns()
}

func (dataset *mixedFeatureTypeDataset) NumFeatures() int {
	return dataset.numFeatures
}

func (dataset *mixedFeatureTypeDataset) AddRowFromStrings(targetStart, targetEnd int, strings []string) error {
	actualLength := len(strings)
	expectedLength := dataset.numColumns()

	if actualLength != expectedLength {
		return newRowLengthMismatchError(actualLength, expectedLength)
	}

	rawValues := make([]float64, actualLength)

	for i, s := range strings {
		columnType := dataset.columnTypes[i]

		value, err := columnType.PersistRawFromString(s)
		if err != nil {
			return err
		}

		rawValues[i] = value
	}

	dataset.rawDataset = append(dataset.rawDataset, rawValues...)

	return nil
}

func (dataset *mixedFeatureTypeDataset) Row(i int) (row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	return inmemoryrow.NewRow(dataset.targetAtRow(i), dataset.rawFeatureValuesAtRow(i), false), nil
}

func (dataset *floatFeatureTypeDataset) FloatFeatureRow(i int) (row.FloatFeatureRow, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	return inmemoryrow.NewRow(
		dataset.targetAtRow(i),
		dataset.rawFeatureValuesAtRow(i),
		true,
	).(row.FloatFeatureRow), nil
}

func (dataset *mixedFeatureTypeDataset) targetAtRow(i int) target.Target {
	numColumns := dataset.numColumns()
	target := []interface{}{}

	for j := dataset.targetStart; j <= dataset.targetEnd; j++ {
		value, _ := dataset.columnTypes[j].ValueFromRaw(dataset.rawDataset[i*numColumns+j])
		target = append(target, value)
	}

	return target
}

func (dataset *mixedFeatureTypeDataset) rawFeatureValuesAtRow(i int) []float64 {
	numColumns := dataset.numColumns()
	rawFeatureValues := dataset.rawDataset[i*numColumns : i*numColumns+dataset.targetStart]

	if dataset.targetEnd < numColumns {
		rawFeatureValues = append(rawFeatureValues, dataset.rawDataset[i*numColumns+dataset.targetEnd+1:(i+1)*numColumns]...)
	}

	return rawFeatureValues
}

func (dataset *mixedFeatureTypeDataset) numColumns() int {
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
