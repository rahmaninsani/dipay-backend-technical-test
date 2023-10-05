package test

import (
	"github.com/rahmaninsani/dipay-backend-technical-test/01-array-operation/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DuplicateAndShiftTest struct {
	arr      []int
	expected []int
}

func TestDuplicateAndShiftSuccess(t *testing.T) {
	testCases := []DuplicateAndShiftTest{
		{
			arr:      []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			arr:      []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		result := helper.DuplicateAndShift(testCase.arr)
		assert.Equal(t, testCase.expected, result)
	}
}

func TestDuplicateAndShiftFailure(t *testing.T) {
	testCases := []DuplicateAndShiftTest{
		{
			arr:      []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4, 5, 0, 0},
		},
		{
			arr:      []int{1, 2, 3},
			expected: []int{1, 2, 3, 0},
		},
	}

	for _, testCase := range testCases {
		result := helper.DuplicateAndShift(testCase.arr)
		assert.NotEqual(t, testCase.expected, result)
	}
}

func BenchmarkDuplicateAndShift(b *testing.B) {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}

	for i := 0; i < b.N; i++ {
		helper.DuplicateAndShift(arr)
	}
}
