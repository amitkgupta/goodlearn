package csvparse

import (
	"encoding/csv"
	"os"

	"github.com/amitkgupta/goodlearn/data/columntype"
	"github.com/amitkgupta/goodlearn/data/dataset/inmemorydataset"
)

func InMemoryDatasetFromPath(filepath string, targetStart, targetEnd int) (*inmemorydataset.InMemoryDataset, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	_, err = reader.Read()
	firstLine, err := reader.Read()
	if err != nil {
		return nil, err
	}

	columnTypes, err := columntype.StringsToColumnTypes(firstLine)
	if err != nil {
		return nil, err
	}

	newDataset, err := inmemorydataset.NewDataset(targetStart, targetEnd, columnTypes)
	if err != nil {
		return nil, err
	}

	err = newDataset.AddRowFromStrings(targetStart, targetEnd, columnTypes, firstLine)
	if err != nil {
		return nil, err
	}

	for line, err := reader.Read(); err == nil; line, err = reader.Read() {
		err = newDataset.AddRowFromStrings(targetStart, targetEnd, columnTypes, line)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	return newDataset, nil
}
