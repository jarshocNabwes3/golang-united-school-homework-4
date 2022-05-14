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

func testStringEmpties(t *testing.T) {
	testSumInOutErr(t, ``, ``, errorEmptyInput)
	testSumInOutErr(t, ` `, ``, errorEmptyInput)
	testSumInOutErr(t, ` 
		`, ``, errorEmptyInput)
}

func testNotSingleOperations(t *testing.T) {
	testSumInOutErr(t, `1`, ``, errorNotSingleOperation)
	testSumInOutErr(t, `11`, ``, errorNotSingleOperation)
	testSumInOutErr(t, `a`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "76+125 + 852+", ``, errorNotSingleOperation)
	testSumInOutErr(t, ` 1+1
	1+1
		1+1`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "-51+1\n912+158\n"+"			1+1\n", ``, errorNotSingleOperation)
	testSumInOutErr(t, ` 41+33
	1+1
		1+1`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "76+1\n1+1\n"+"			1+1\n", ``, errorNotSingleOperation)

	testSumInOutErr(t, `a1`, ``, errorNotSingleOperation)
	testSumInOutErr(t, `a11`, ``, errorNotSingleOperation)
	testSumInOutErr(t, `1a`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "7a6+12z5 + 8M52+", ``, errorNotSingleOperation)
	testSumInOutErr(t, ` S1+1S
	1N+N1
		abcd1+1efgh`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "-5d1+1\n9S12+15Y8\n"+"			1+1\n", ``, errorNotSingleOperation)
	testSumInOutErr(t, ` 4U1+33q
	P1+1t
		f1+L1`, ``, errorNotSingleOperation)
	testSumInOutErr(t, "7R6+1\n1o+1\n"+"			1+1\n", ``, errorNotSingleOperation)
}

func testNotTwoOperands(t *testing.T) {
	testSumInOutErr(t, "76+125 + 852", ``, errorNotTwoOperands)
	testSumInOutErr(t, "76-125 + 852", ``, errorNotTwoOperands)
	testSumInOutErr(t, "-76+125 - 852", ``, errorNotTwoOperands)

	testSumInOutErr(t, "a76+125z + m8n52", ``, errorNotTwoOperands)
	testSumInOutErr(t, "7p6-q1w2k5v + 852k", ``, errorNotTwoOperands)
	testSumInOutErr(t, "-a76s+f1d25 - 8w5e2C", ``, errorNotTwoOperands)
}

func testSumInOut(t *testing.T, input, resultExpected string) {
	result, err := StringSum(input)

	// test error
	errToTest := errors.Unwrap(err)
	assert.Equal(t, errToTest == nil, true, "Input: '%v' Error: '%v'", input, errToTest)

	// test result
	assert.Equal(t, resultExpected, result, `Result string differs from expected: Input: '%v'; Expected: '%v'; Result: '%v'`, input, resultExpected, result)
}

func testSums(t *testing.T) {
	testSumInOut(t, "-76+125", "49")
	testSumInOut(t, "+76-125", "-49")
	testSumInOut(t, "76+125", "201")

	testSumInOut(t, "-76+125\n", "49")
	testSumInOut(t, "+76-125\n", "-49")
	testSumInOut(t, "76+125\n", "201")
}

func TestStringSum(t *testing.T) {
	testStringEmpties(t)
	testNotSingleOperations(t)
	testNotTwoOperands(t)

	testSums(t)
}
