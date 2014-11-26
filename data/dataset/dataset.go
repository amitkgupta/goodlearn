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
	AllTargetsFloats() bool

	NumFeatures() int
	NumTargets() int

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
	numTargets           int
	numColumns           int
	rows                 []row.Row
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
		len(targetColumnIndices),
		len(columnTypes),
		[]row.Row{},
	}
}

func (dataset *inMemoryDataset) AllFeaturesFloats() bool {
	return dataset.allFeaturesFloats
}

func (dataset *inMemoryDataset) AllTargetsFloats() bool {
	return dataset.allTargetsFloats
}

func (dataset *inMemoryDataset) NumFeatures() int {
	return dataset.numFeatures
}

func (dataset *inMemoryDataset) NumTargets() int {
	return dataset.numTargets
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

	features, _ := slice.SliceFromRawValues(
		dataset.allFeaturesFloats,
		dataset.featureColumnIndices,
		dataset.columnTypes,
		rawValues,
	)

	target, _ := slice.SliceFromRawValues(
		dataset.allTargetsFloats,
		dataset.targetColumnIndices,
		dataset.columnTypes,
		rawValues,
	)

	dataset.rows = append(dataset.rows, row.NewRow(features, target, dataset.numFeatures))

	return nil
}

func (dataset *inMemoryDataset) NumRows() int {
	return len(dataset.rows)
}

func (dataset *inMemoryDataset) Row(i int) (row.Row, error) {
	numRows := dataset.NumRows()
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	return dataset.rows[i], nil
}

func NewSubset(ds Dataset, rowMap []int) Dataset {
	return &subset{
		ds,
		rowMap,
		ds.AllFeaturesFloats(),
		ds.AllTargetsFloats(),
		ds.NumFeatures(),
		ds.NumTargets(),
		len(rowMap),
	}
}

type subset struct {
	superset          Dataset
	rowMap            []int
	allFeaturesFloats bool
	allTargetsFloats  bool
	numFeatures       int
	numTargets        int
	numRows           int
}

func (s *subset) AllFeaturesFloats() bool {
	return s.allFeaturesFloats
}

func (s *subset) AllTargetsFloats() bool {
	return s.allTargetsFloats
}

func (s *subset) NumFeatures() int {
	return s.numFeatures
}

func (s *subset) NumTargets() int {
	return s.numTargets
}

func (s *subset) AddRowFromStrings([]string) error {
	return errors.New("AddRowFromStrings operation not permitted on subsets")
}

func (s *subset) NumRows() int {
	return s.numRows
}

func (s *subset) Row(i int) (row.Row, error) {
	numRows := s.numRows
	if i < 0 || numRows <= i {
		return nil, newDatasetRowIndexOutOfBoundsError(i, numRows)
	}

	return s.superset.Row(s.rowMap[i])
}

func newRowLengthMismatchError(actual, expected int) error {
	return errors.New(fmt.Sprintf("Row has length %d, expected %d", actual, expected))
}

func newDatasetRowIndexOutOfBoundsError(index, numRows int) error {
	return errors.New(fmt.Sprintf("Cannot access row %d in dataset with %d rows", index, numRows))
}
