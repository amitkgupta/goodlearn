package columntype

import (
	"strconv"
)

type floatType struct{}

func newFloatType() *floatType {
	return &floatType{}
}

func (ft *floatType) ValueFromRaw(x float64) (interface{}, error) {
	return x, nil
}

func (ft *floatType) PersistRawFromString(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (ft *floatType) IsFloat() bool {
	return true
}
