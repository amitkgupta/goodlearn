package columntype

import (
	"strconv"

	"github.com/amitkgupta/goodlearn/data/columntype/floatcolumntype"
	"github.com/amitkgupta/goodlearn/data/columntype/stringcolumntype"
)

type ColumnType interface {
	PersistRawFromString(string) (float64, error)
	ValueFromRaw(float64) (interface{}, error)
	IsFloat() bool
}

func StringsToColumnTypes(strings []string) ([]ColumnType, error) {
	types := make([]ColumnType, len(strings))

	for i, s := range strings {
		_, err := strconv.ParseFloat(s, 64)

		if err == nil {
			types[i] = floatcolumntype.NewFloatType()
		} else {
			if err.(*strconv.NumError).Err == strconv.ErrSyntax {
				types[i] = stringcolumntype.NewStringType()
			} else {
				return nil, err
			}
		}
	}

	return types, nil
}
