package row

import (
	"github.com/amitkgupta/goodlearn/dataset/target"
)

type Row struct {
	Target            target.Target
	rawFeatureValues  []float64
	AllFeaturesFloats bool
	NumFeatures       int
}

func UnsafeNewRow(target target.Target, rawFeatureValues []float64, allFeaturesFloats bool, numFeatures int) *Row {
	return &Row{target, rawFeatureValues, allFeaturesFloats, numFeatures}
}

func (r *Row) UnsafeFloatFeatureValues() []float64 {
	return r.rawFeatureValues
}
