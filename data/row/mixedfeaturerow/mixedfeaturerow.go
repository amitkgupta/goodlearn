package mixedfeaturerow

import (
	"github.com/amitkgupta/goodlearn/data/target"
)

type mixedFeatureRow struct {
	target      target.Target
	numFeatures int
}

func NewMixedFeatureRow(target target.Target, rawFeatureValues []float64) *mixedFeatureRow {
	return &mixedFeatureRow{target, len(rawFeatureValues)}
}

func (row *mixedFeatureRow) Target() target.Target {
	return row.target
}

func (row *mixedFeatureRow) NumFeatures() int {
	return row.numFeatures
}

func (row *mixedFeatureRow) Features() []interface{} {
	panic("not implemented")
	return nil
}
