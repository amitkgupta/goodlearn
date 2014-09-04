package slice

import (
	"errors"
	"fmt"

	"github.com/amitkgupta/goodlearn/data/columntype"
)

type Slice interface {
	len() int
	entry(int) interface{}
	Equals(Slice) bool
}

type FloatSlice interface {
	Slice
	Values() []float64
}

type MixedSlice interface {
	Slice
	Values() []interface{}
}

type floatSlice struct {
	values []float64
}

type mixedSlice struct {
	values []interface{}
}

func SliceFromRawValues(
	allFloats bool,
	columnIndices []int,
	columnTypes []columntype.ColumnType,
	rawValues []float64,
) (Slice, error) {
	var err error

	if allFloats {
		values := make([]float64, len(columnIndices))

		for idx, i := range columnIndices {
			if floatColumnType, ok := columnTypes[i].(columntype.FloatColumnType); ok {
				values[idx], err = floatColumnType.ValueFromRaw(rawValues[i])
				if err != nil {
					return nil, err
				}
			} else {
				return nil, newExpectedAllColumnsFloatsError(i, columnIndices)
			}
		}

		return &floatSlice{values}, nil
	} else {
		values := make([]interface{}, len(columnIndices))

		for idx, i := range columnIndices {
			if floatColumnType, ok := columnTypes[i].(columntype.FloatColumnType); ok {
				values[idx], err = floatColumnType.ValueFromRaw(rawValues[i])
				if err != nil {
					return nil, err
				}
			} else if stringColumnType, ok := columnTypes[i].(columntype.StringColumnType); ok {
				values[idx], err = stringColumnType.ValueFromRaw(rawValues[i])
				if err != nil {
					return nil, err
				}
			} else {
				return nil, newUnknownColumnTypeError()
			}
		}

		return &mixedSlice{values}, nil
	}
}

func (s *floatSlice) len() int {
	return len(s.values)
}

func (s *mixedSlice) len() int {
	return len(s.values)
}

func (s *floatSlice) entry(i int) interface{} {
	return s.values[i]
}

func (s *mixedSlice) entry(i int) interface{} {
	return s.values[i]
}

func (s *floatSlice) Equals(other Slice) bool {
	return compare(s, other)
}

func (s *mixedSlice) Equals(other Slice) bool {
	return compare(s, other)
}

func compare(s1, s2 Slice) bool {
	if s1.len() != s2.len() {
		return false
	}

	for i := 0; i < s1.len(); i++ {
		if s1.entry(i) != s2.entry(i) {
			return false
		}
	}

	return true
}

func (s *floatSlice) Values() []float64 {
	return s.values
}

func (s *mixedSlice) Values() []interface{} {
	return s.values
}

func newUnknownColumnTypeError() error {
	return errors.New("Unknown column type error")
}

func newExpectedAllColumnsFloatsError(i int, columnIndices []int) error {
	return errors.New(fmt.Sprintf(
		"Column %d is not a float column, expected all columns in %v to be float columns",
		i,
		columnIndices,
	))
}
