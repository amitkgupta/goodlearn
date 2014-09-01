package columntype

import (
	"strconv"

	"github.com/amitkgupta/goodlearn/data/columntype/floatcolumntype"
	"github.com/amitkgupta/goodlearn/data/columntype/stringcolumntype"
)

type ColumnType interface {
	PersistRawFromString(string) (float64, error)
}

type FloatColumnType interface {
	ColumnType
	ValueFromRaw(float64) (float64, error)
}

type StringColumnType interface {
	ColumnType
	ValueFromRaw(float64) (string, error)
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
