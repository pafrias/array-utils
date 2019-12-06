package arr

import (
	"errors"
	"fmt"
	"reflect"
)

/*matchFuncType tests for matching elements in a slice and in a functions input, for use as
a predicate in slice utility functions.

It also tests if that function returns bool*/
func matchFuncType(src, test reflect.Type) error {
	var errorStr string

	sourceKind := src.Kind()
	testKind := test.Kind()
	if sourceKind != reflect.Slice {
		errorStr = fmt.Sprintf("Expected type (slice) but received src of type (%v)\n", sourceKind)
	} else if testKind != reflect.Func {
		errorStr = fmt.Sprintf("Expected type (func) but received testFunc of type (%v)\n", testKind)
	}

	testOutput := test.Out(0).Kind()
	testInput := test.In(0)
	elemType := src.Elem()
	if test.NumIn() != 1 {
		errorStr = "Expected testFunc to have 1 input\n"
	} else if testInput != elemType {
		errorStr = fmt.Sprintf("Received slice of type (%v), but testFunc expects type (%v)", elemType, testInput)
	} else if test.NumOut() != 1 || testOutput != reflect.Bool {
		errorStr = "Expected func to have 1 return of type (bool)\n"
	}

	if errorStr != "" {
		return errors.New(errorStr)
	}

	return nil
}

func matchSliceType(src, test reflect.Type) (reflect.Type, bool) {
	sk := src.Kind()
	if sk != reflect.Slice {
		return nil, false
	}
	elemType := src.Elem()
	return elemType, elemType == test
}
