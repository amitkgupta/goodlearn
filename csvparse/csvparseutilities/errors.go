package csvparseutilities

import (
	"fmt"
)

func NewUnableToOpenFileError(filepath string, err error) UnableToOpenFileError {
	return UnableToOpenFileError{filepath, err}
}
func NewUnableToReadTwoLinesError(filepath string, err error) UnableToReadTwoLinesError {
	return UnableToReadTwoLinesError{filepath, err}
}
func NewUnableToParseColumnTypesError(filepath string, err error) UnableToParseColumnTypesError {
	return UnableToParseColumnTypesError{filepath, err}
}

func NewTargetOutOfBoundsError(filepath string, targetStartInclusive, targetEndExclusive, numColumns int) error {
	return TargetOutOfBoundsError{
		filepath,
		targetStartInclusive,
		targetEndExclusive,
		numColumns,
	}
}
func NewUnableToParseRowError(filepath string, err error) UnableToParseRowError {
	return UnableToParseRowError{filepath, err}
}
func NewGenericError(filepath string, err error) GenericError {
	return GenericError{filepath, err}
}

type baseError struct {
	filepath string
	err      error
}
type UnableToOpenFileError baseError
type UnableToReadTwoLinesError baseError
type UnableToParseColumnTypesError baseError
type TargetOutOfBoundsError struct {
	filepath             string
	targetStartInclusive int
	targetEndExclusive   int
	numColumns           int
}
type UnableToParseRowError baseError
type GenericError baseError

func (e UnableToOpenFileError) Error() string {
	return fmt.Sprintf("Unable to open file at '%s': %s", e.filepath, e.err.Error())
}
func (e UnableToReadTwoLinesError) Error() string {
	return fmt.Sprintf("Unable to read at least two lines from '%s': %s", e.filepath, e.err.Error())
}
func (e UnableToParseColumnTypesError) Error() string {
	return fmt.Sprintf("Unable to parse column types for '%s': %s", e.filepath, e.err.Error())
}
func (e TargetOutOfBoundsError) Error() string {
	return fmt.Sprintf(
		"Unable to create dataset from '%s'; columns must have valid target bounds, and at least one non-target column; "+
			"cannot have %d total columns, target start column %d and target end column %d",
		e.filepath,
		e.numColumns,
		e.targetStartInclusive,
		e.targetEndExclusive,
	)

}
func (e UnableToParseRowError) Error() string {
	return fmt.Sprintf("Unable to parse some row in '%s': %s", e.filepath, e.err.Error())
}
func (e GenericError) Error() string {
	return fmt.Sprintf("An error occurred parsing '%s' to a dataset: %s", e.filepath, e.err.Error())
}
