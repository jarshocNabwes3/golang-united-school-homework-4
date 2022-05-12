package string_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testEmptyResultAndError(t *testing.T, input, resultExpected string, errExpected error) {
	result, err := StringSum(``)
	assert.EqualErrorf(t, err, errorEmptyInput.Error(), "Input: '%v' Error should be: %v, got: %v", input, errExpected.Error(), err)
	assert.Equal(t, ``, result, `Result string differs from expected: Input: '%v'; Expected: '%v'; Result: '%v'`, ``, resultExpected, result)
}

func TestStringSum(t *testing.T) {
	testEmptyResultAndError(t, ``, ``, errorEmptyInput)
	testEmptyResultAndError(t, ` `, ``, errorEmptyInput)
	testEmptyResultAndError(t, ` 
		``, ``, errorEmptyInput)
}
