package row

import (
	"github.com/amitkgupta/goodlearn/dataset/columntype"
	"github.com/amitkgupta/goodlearn/dataset/target"
)

type Row struct {
	Target           target.Target
	rawFeatureValues []float64
}

// bad, handle error
func NewRow(rawValues []float64, targetStart, targetEnd int, columnTypes []columntype.ColumnType) (*Row, error) {
	target := []interface{}{}
	for i, rawValue := range rawValues {
		value, err := columnTypes[i].ValueFromRaw(rawValue)
		if err != nil {
			return nil, err
		}

		if targetStart <= i && i <= targetEnd {
			target = append(target, value)
		}
	}

	rawFeatureValues := rawValues[0:targetStart]
	if targetEnd < len(columnTypes)-1 {
		rawFeatureValues = append(rawFeatureValues[targetEnd+1:])
	}

	return &Row{target, rawFeatureValues}, nil
}

func (r *Row) UnsafeFloatFeatureValues() []float64 {
	return r.rawFeatureValues
}
