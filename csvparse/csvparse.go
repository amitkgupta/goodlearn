package csvparse

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/amitkgupta/goodlearn/csvparse/csvparseutilities"
	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset"
)

func DatasetFromPath(filepath string, targetStartInclusive, targetEndExclusive int) (dataset.Dataset, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, csvparseutilities.NewUnableToOpenFileError(filepath, err)
	}

	reader := csv.NewReader(file)

	_, err = reader.Read()
	line, err := reader.Read()
	if err != nil {
		return nil, csvparseutilities.NewUnableToReadTwoLinesError(filepath, err)
	}

	columnTypes, err := columntype.StringsToColumnTypes(line)
	if err != nil {
		return nil, csvparseutilities.NewUnableToParseColumnTypesError(filepath, err)
	}

	newDataset, err := dataset.NewDataset(targetStartInclusive, targetEndExclusive, columnTypes)
	if err != nil {
		return nil, csvparseutilities.NewUnableToCreateDatasetError(filepath, err)
	}

	for ; err == nil; line, err = reader.Read() {
		err = newDataset.AddRowFromStrings(line)
		if err != nil {
			return nil, csvparseutilities.NewUnableToParseRowError(filepath, err)
		}
	}
	if err != nil && err != io.EOF {
		return nil, csvparseutilities.NewGenericError(filepath, err)
	}

	return newDataset, nil
}
