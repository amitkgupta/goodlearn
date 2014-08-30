package columntype

import (
	"strconv"
)

type ColumnType interface {
	RawFromString(string) (float64, error)
	ValueFromRaw(float64) (interface{}, error)
}

func StringsToColumnTypes(strings []string) ([]ColumnType, error) {
	types := make([]ColumnType, len(strings))

	for i, s := range strings {
		_, err := strconv.ParseFloat(s, 64)

		if err == nil {
			types[i] = newFloatType()
		} else {
			if err.(*strconv.NumError).Err == strconv.ErrSyntax {
				types[i] = newStringType()
			} else {
				return nil, err
			}
		}
	}

	return types, nil
}
