package columntype

import (
	"errors"
	"fmt"
	"strconv"
)

type ColumnType interface {
	PersistRawFromString(string) (float64, error)
}

type FloatColumnType interface {
	ColumnType
	ValueFromRaw(float64) float64
}

type StringColumnType interface {
	ColumnType
	ValueFromRaw(float64) (string, error)
}

type floatType struct{}

type stringType struct {
	counter  float64
	encoding map[string]float64
	decoding map[float64]string
}

func StringsToColumnTypes(strings []string) ([]ColumnType, error) {
	types := make([]ColumnType, len(strings))

	for i, s := range strings {
		_, err := strconv.ParseFloat(s, 64)

		if err == nil {
			types[i] = &floatType{}
		} else {
			if err.(*strconv.NumError).Err == strconv.ErrSyntax {
				types[i] = &stringType{
					0,
					make(map[string]float64),
					make(map[float64]string),
				}
			} else {
				return nil, newUnableToParseLargeFloatError(s)
			}
		}
	}

	return types, nil
}

func (ft *floatType) ValueFromRaw(x float64) float64 {
	return x
}

func (ft *floatType) PersistRawFromString(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (st *stringType) ValueFromRaw(raw float64) (string, error) {
	value, ok := st.decoding[raw]
	if !ok {
		return "", newUnknownCodeError(raw)
	}

	return value, nil
}

func (st *stringType) PersistRawFromString(s string) (float64, error) {
	value, ok := st.encoding[s]
	if ok {
		return value, nil
	}

	st.encoding[s] = st.counter
	st.decoding[st.counter] = s
	st.counter++

	return st.encoding[s], nil
}

func newUnknownCodeError(raw float64) error {
	return errors.New(fmt.Sprintf("Unknown code %v", raw))
}

func newUnableToParseLargeFloatError(s string) error {
	return errors.New(fmt.Sprintf("Unable to parse '%s' into 64-bit float"))
}
