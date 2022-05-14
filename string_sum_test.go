package string_sum

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSumInOutErr(t *testing.T, input, resultExpected string, errExpected error) {
	result, err := StringSum(input)

	// test error
	errToTest := errors.Unwrap(err)
	errStringExpected := errExpected.Error()
	assert.EqualErrorf(t, errToTest, errStringExpected, "Input: '%v' Error should be: '%v', got: '%v'", input, errStringExpected, errToTest)

	// test result
	assert.Equal(t, resultExpected, result, `Result string differs from expected: Input: '%v'; Expected: '%v'; Result: '%v'`, input, resultExpected, result)
}

func TestStringSum(t *testing.T) {
	testSumInOutErr(t, ``, ``, errorEmptyInput)
	testSumInOutErr(t, ` `, ``, errorEmptyInput)
	testSumInOutErr(t, ` 
		`, ``, errorEmptyInput)
}
