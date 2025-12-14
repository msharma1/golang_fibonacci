package main

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

type fibTest struct {
	input    int
	expected []int
	hasError bool
}

func TestFibonacci(t *testing.T) {
	tests := []fibTest{
		// Basic valid cases
		{input: 0, expected: []int{}, hasError: false},
		{input: 1, expected: []int{0}, hasError: false},
		{input: 2, expected: []int{0, 1}, hasError: false},
		{input: 5, expected: []int{0, 1, 1, 2, 3}, hasError: false},
		{input: 10, expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, hasError: false},

		{input: -1, expected: nil, hasError: true},
	}

	for _, test := range tests {
		actual, err := generateFibonacci(test.input)

		if test.hasError {
			assert.Error(t, err, "Input %d should return an error", test.input)
		} else {
			assert.NoError(t, err, "Input %d should not return an error", test.input)

			match := slices.Equal(actual, test.expected)
			assert.True(t, match, "Input %d mismatch! Expected %v, got %v", test.input, test.expected, actual)
		}
	}
}
