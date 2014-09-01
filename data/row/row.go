package row

import (
	"github.com/amitkgupta/goodlearn/data/target"
)

type Row interface {
	Target() target.Target
	NumFeatures() int
}

type FloatFeatureRow interface {
	Row
	FloatFeatureValues() []float64
}
