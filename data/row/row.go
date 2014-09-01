package row

import (
	"github.com/amitkgupta/goodlearn/data/target"
)

type Row struct {
	Target            target.Target
	rawFeatureValues  []float64
	allFeaturesFloats bool
	numFeatures       int
}

func UnsafeNewRow(target target.Target, rawFeatureValues []float64, allFeaturesFloats bool) *Row {
	return &Row{target, rawFeatureValues, allFeaturesFloats, len(rawFeatureValues)}
}

func (r *Row) AllFeaturesFloats() bool {
	return r.allFeaturesFloats
}

func (r *Row) NumFeatures() int {
	return r.numFeatures
}

func (r *Row) UnsafeFloatFeatureValues() []float64 {
	return r.rawFeatureValues
}
