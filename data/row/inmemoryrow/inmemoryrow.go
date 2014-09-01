package inmemoryrow

import (
	"github.com/amitkgupta/goodlearn/data/row"
	"github.com/amitkgupta/goodlearn/data/target"
)

type mixedFeatureTypeRow struct {
	target           target.Target
	rawFeatureValues []float64
	numFeatures      int
}

type floatFeatureTypeRow struct {
	*mixedFeatureTypeRow
}

func NewRow(target target.Target, rawFeatureValues []float64, allFeaturesFloat bool) row.Row {
	mixedFeatureTypeRow := &mixedFeatureTypeRow{target, rawFeatureValues, len(rawFeatureValues)}

	if allFeaturesFloat {
		return &floatFeatureTypeRow{mixedFeatureTypeRow}
	} else {
		return mixedFeatureTypeRow
	}
}

func (r *mixedFeatureTypeRow) Target() target.Target {
	return r.target
}

func (r *mixedFeatureTypeRow) NumFeatures() int {
	return r.numFeatures
}

func (r *floatFeatureTypeRow) FloatFeatureValues() []float64 {
	return r.rawFeatureValues
}
