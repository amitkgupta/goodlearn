package stringcolumntype

import (
	"errors"
	"fmt"
)

type stringType struct {
	counter  float64
	encoding map[string]float64
	decoding map[float64]string
}

func NewStringType() *stringType {
	return &stringType{0, make(map[string]float64), make(map[float64]string)}
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