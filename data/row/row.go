package row

import (
	"github.com/amitkgupta/goodlearn/data/slice"
)

type Row interface {
	Features() slice.Slice
	Target() slice.Slice
	NumFeatures() int
}

type row struct {
	features    slice.Slice
	target      slice.Slice
	numFeatures int
}

func NewRow(features, target slice.Slice, numFeatures int) Row {
	return &row{features, target, numFeatures}
}

func (r *row) Features() slice.Slice {
	return r.features
}

func (r *row) Target() slice.Slice {
	return r.target
}

func (r *row) NumFeatures() int {
	return r.numFeatures
}
