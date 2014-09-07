package csvparse

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
	"github.com/amitkgupta/goodlearn/errors/csvparseerrors"
)

func DatasetFromPath(filepath string, targetStartInclusive, targetEndExclusive int) (dataset.Dataset, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, csvparseerrors.NewUnableToOpenFileError(filepath, err)
	}

	reader := csv.NewReader(file)

	_, err = reader.Read()
	line, err := reader.Read()
	if err != nil {
		return nil, csvparseerrors.NewUnableToReadTwoLinesError(filepath, err)
	}

	columnTypes, err := columntype.StringsToColumnTypes(line)
	if err != nil {
		return nil, csvparseerrors.NewUnableToParseColumnTypesError(filepath, err)
	}

	numColumns := len(columnTypes)
	if targetOutOfBounds(targetStartInclusive, targetEndExclusive, numColumns) {
		return nil, csvparseerrors.NewTargetOutOfBoundsError(filepath, targetStartInclusive, targetEndExclusive, numColumns)
	}

	newDataset := dataset.NewDataset(
		featureColumnIndices(targetStartInclusive, targetEndExclusive, numColumns),
		targetColumnIndices(targetStartInclusive, targetEndExclusive, numColumns),
		columnTypes,
	)

	for ; err == nil; line, err = reader.Read() {
		err = newDataset.AddRowFromStrings(line)
		if err != nil {
			return nil, csvparseerrors.NewUnableToParseRowError(filepath, err)
		}
	}
	if err != nil && err != io.EOF {
		return nil, csvparseerrors.NewGenericError(filepath, err)
	}

	return newDataset, nil
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
