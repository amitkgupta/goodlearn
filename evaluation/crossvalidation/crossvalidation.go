package crossvalidation

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/amitkgupta/goodlearn/data/dataset"
)

func SplitDataset(ds dataset.Dataset, trainingRatio float64, source rand.Source) (dataset.Dataset, dataset.Dataset, error) {
	if trainingRatio < 0 || trainingRatio > 1 {
		return nil, nil, fmt.Errorf("Unable to split dataset with invalid ratio %.2f", trainingRatio)
	}

	numRows := ds.NumRows()
	if numRows == 0 {
		return nil, nil, errors.New("Cannot split empty dataset")
	}

	r := rand.New(source)
	perm := r.Perm(numRows)

	trainingRowMap := make([]int, 0, numRows)
	testRowMap := make([]int, 0, numRows)

	for _, rowIndex := range perm {
		if r.Float64() < trainingRatio {
			trainingRowMap = append(trainingRowMap, rowIndex)
		} else {
			testRowMap = append(testRowMap, rowIndex)
		}
	}

	return dataset.NewSubset(ds, trainingRowMap), dataset.NewSubset(ds, testRowMap), nil
}
