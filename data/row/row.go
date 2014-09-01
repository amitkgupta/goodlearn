package row

import (
	"github.com/amitkgupta/goodlearn/data/row/floatfeaturerow"
	"github.com/amitkgupta/goodlearn/data/row/mixedfeaturerow"
	"github.com/amitkgupta/goodlearn/data/target"
)

type Row interface {
	Target() target.Target
	NumFeatures() int
}

type MixedFeatureRow interface {
	Row
	Features() []interface{}
}

type FloatFeatureRow interface {
	Row
	Features() []float64
}

func NewRow(allFeaturesFloats bool, target target.Target, rawFeatureValues []float64) Row {
	if allFeaturesFloats {
		return floatfeaturerow.NewFloatFeatureRow(target, rawFeatureValues)
	} else {
		return mixedfeaturerow.NewMixedFeatureRow(target, rawFeatureValues)
	}
}
