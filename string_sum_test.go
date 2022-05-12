package string_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSum(t *testing.T) {
	result, err := StringSum(``)
	assert.EqualErrorf(t, err, errorEmptyInput.Error(), "Error should be: %v, got: %v", errorEmptyInput.Error(), err)
	assert.Equal(t, ``, result, `Result string differs from expected: Input: '%v'; Expected: '%v'; Result: '%v'`, ``, ``, result)
}
