package string_sum

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands     = errors.New("expecting two operands, but received more or less")
	errorNotSingleOperation = errors.New("expecting one operation, but received")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func findOperation(input string) (operation []string, err error) {

	input = regexp.MustCompile(`(?ms)[ \t\r]+`).ReplaceAllString(input, ``)
	inputSingleLine := regexp.MustCompile(`\n`).ReplaceAllString(input, ``)
	if inputSingleLine == `` {
		return make([]string, 0), fmt.Errorf(`sum: %w`, errorEmptyInput)
	}

	input = regexp.MustCompile(`(?s)([^\n])$`).ReplaceAllString(input,
		"$1\n")

	operations := regexp.MustCompile(`([-+]?[^-+\s]+)([-+][^-+\s]+)([-+][^-+\s]+)?[$\n]`).FindAllStringSubmatch(input, 2)
	if operations == nil {
		return make([]string, 0), fmt.Errorf(`sum: %w: '%v'`, errorNotSingleOperation, 0)
	}
	operationsCount := len(operations)
	if operationsCount != 1 {
		return make([]string, 0), fmt.Errorf(`sum: %w: '%v'`, errorNotSingleOperation, operationsCount)
	}
	operation = operations[0]

	return
}

func StringSum(input string) (output string, err error) {
	operation, err := findOperation(input)
	if err != nil {
		return ``, err
	}

	operationCount := len(operation)
	if operationCount != 4 {
		return ``, fmt.Errorf(`sum: %w`, errorNotTwoOperands)
	}
	lastOperand := operation[3:4][0]
	if lastOperand != `` {
		return ``, fmt.Errorf(`sum: %w`, errorNotTwoOperands)
	}
	operands := operation[1:3]

	sum := 0
	for i := range operands {
		operand := 0

		operandString := operands[i]
		operand, err = strconv.Atoi(operandString)
		if err != nil {
			err = fmt.Errorf(`sum: strconv Atoi error: '%w', operand: '%v'`, err, operandString)
			break
		}

		sum += operand
	}

	if err != nil {
		return ``, err
	}

	output = strconv.Itoa(sum)

	return
}
