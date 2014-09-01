package floatfeaturerow

import (
	"github.com/amitkgupta/goodlearn/data/target"
)

type floatFeatureRow struct {
	target           target.Target
	numFeatures      int
	rawFeatureValues []float64
}

func NewFloatFeatureRow(target target.Target, rawFeatureValues []float64) *floatFeatureRow {
	return &floatFeatureRow{target, len(rawFeatureValues), rawFeatureValues}
}

func (row *floatFeatureRow) Target() target.Target {
	return row.target
}

func (row *floatFeatureRow) NumFeatures() int {
	return row.numFeatures
}

func (row *floatFeatureRow) Features() []float64 {
	return row.rawFeatureValues
}
