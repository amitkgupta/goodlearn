package dataset

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/dataset/columntype"
	"github.com/amitkgupta/goodlearn/dataset/row"
)

type Dataset struct {
	rawDataset  []float64
	targetStart int
	targetEnd   int
	columnTypes []columntype.ColumnType
}

func NewDataset(targetStart, targetEnd int, columnTypes []columntype.ColumnType) (*Dataset, error) {
	numColumns := len(columnTypes)

	if targetOutOfBounds(targetStart, targetEnd, numColumns) {
		return nil, newTargetOutOfBoundsError(targetStart, targetEnd, numColumns)
	}

	return &Dataset{[]float64{}, targetStart, targetEnd, columnTypes}, nil
}

func (dataset *Dataset) NumRows() int {
	return len(dataset.rawDataset) / dataset.numColumns()
}

func (dataset *Dataset) Row(i int) (*row.Row, error) {
	numRows := dataset.NumRows()

	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	numColumns := dataset.numColumns()
	rawValues := dataset.rawDataset[i*numColumns : (i+1)*numColumns]
	return row.NewRow(rawValues, dataset.targetStart, dataset.targetEnd, dataset.columnTypes)
}

//bad, check compatibility
func (dataset *Dataset) AddRowFromStrings(targetStart, targetEnd int, columnTypes []columntype.ColumnType, strings []string) error {
	actualLength := len(strings)
	expectedLength := len(columnTypes)

	if actualLength != expectedLength {
		return newRowLengthMismatchError(actualLength, expectedLength)
	}

	vs := make([]float64, actualLength)

	for i, s := range strings {
		columnType := columnTypes[i]

		v, err := columnType.RawFromString(s)
		if err != nil {
			return err
		}

		vs[i] = v
	}

	dataset.rawDataset = append(dataset.rawDataset, vs...)
	return nil
}

func (dataset *Dataset) numColumns() int {
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
