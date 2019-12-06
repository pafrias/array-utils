package slice

import (
	"errors"
	"fmt"
	"reflect"
)

/*Any traverses the src slice, invoking the predicate function on all elements.

If any invocation returns true, return the index of first occurance.
Else, return -1 (and error is applicable).*/
func Any(src, testFunc interface{}) (int, error) {
	err := matchFuncType(reflect.TypeOf(src), reflect.TypeOf(testFunc))
	if err != nil {
		fmt.Println(err)
	} else {
		srcVal := reflect.ValueOf(src)
		funcVal := reflect.ValueOf(testFunc)

		var input [1]reflect.Value
		for i := 0; i < srcVal.Len(); i++ {
			input[0] = srcVal.Index(i)
			b := funcVal.Call(input[:])[0]
			if b.Bool() {
				return i, nil
			}
		}
	}
	return -1, err
}

/*Every traverses the src slice, invoking the predicate function on all elements.

If any invocation returns false, return false.
if all return true, return true.
Or else if any error occurs, return false and the error.*/
func Every(src, testFunc interface{}) (bool, error) {
	err := matchFuncType(reflect.TypeOf(src), reflect.TypeOf(testFunc))
	if err != nil {
		fmt.Println(err)
	} else {
		srcVal := reflect.ValueOf(src)
		funcVal := reflect.ValueOf(testFunc)

		var input [1]reflect.Value
		for i := 0; i < srcVal.Len(); i++ {
			input[0] = srcVal.Index(i)
			b := funcVal.Call(input[:])[0]
			if !b.Bool() {
				return false, nil
			}
		}
	}
	return true, err
}

/*Filter accepts an interface slice and predicate function, calls the predicate function on every element, and
returns a new slice for all elements for which the predicate returned true.

Bad functions and mismatching values/parameters return nil, empty arrays return a new empty array*/
func Filter(src, testFunc interface{}) ([]interface{}, error) {
	var results = []interface{}{}

	err := matchFuncType(reflect.TypeOf(src), reflect.TypeOf(testFunc))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	srcVal := reflect.ValueOf(src)
	funcVal := reflect.ValueOf(testFunc)

	var input [1]reflect.Value
	for i := 0; i < srcVal.Len(); i++ {
		input[0] = srcVal.Index(i)
		b := funcVal.Call(input[:])[0]
		if b.Bool() {
			in := input[0].Interface()
			results = append(results, in)
		}
	}
	return results, nil
}

/*Has traverses the src slice, invoking the predicate function on all elements.

If any invocation returns true, return the index of first occurance.
Else, return -1 (and error is applicable).*/
func Has(src, target interface{}) (bool, error) {

	_, ok := matchSliceType(reflect.TypeOf(src), reflect.TypeOf(target))
	if !ok {
		return false, errors.New("Cannot search for element in slice of wrong type")
	}

	srcVal := reflect.ValueOf(src)
	tarVal := reflect.ValueOf(target)

	for i := 0; i < srcVal.Len(); i++ {
		if tarVal == srcVal.Index(i) {
			return true, nil
		}
	}
	return false, nil
}
