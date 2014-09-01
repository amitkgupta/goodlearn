package floatcolumntype

import (
	"strconv"
)

type floatType struct{}

func NewFloatType() *floatType {
	return &floatType{}
}

func (ft *floatType) ValueFromRaw(x float64) (float64, error) {
	return x, nil
}

func (ft *floatType) PersistRawFromString(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
