package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleSum(t *testing.T) {
	tcs := map[string]struct {
		value1, value2 float64
		expectedResult float64
	}{
		"add_1": {
			value1: 1, value2: 2,
			expectedResult: 3,
		},
		"add_2": {
			value1: 2, value2: 2,
			expectedResult: 4,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			result := Add(tc.value1, tc.value2)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestExampleSubtract(t *testing.T) {
	tcs := map[string]struct {
		value1, value2 float64
		expectedResult float64
	}{
		"sub_1": {
			value1: 1, value2: 2,
			expectedResult: -1,
		},
		"sub_2": {
			value1: 2, value2: 2,
			expectedResult: 0,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			result := Subtract(tc.value1, tc.value2)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestExampleTimes(t *testing.T) {
	tcs := map[string]struct {
		value1, value2 float64
		expectedResult float64
	}{
		"sub_1": {
			value1: 1, value2: 2,
			expectedResult: 2,
		},
		"sub_2": {
			value1: 2, value2: 2,
			expectedResult: 4,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			result := Times(tc.value1, tc.value2)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestExampleDivide(t *testing.T) {
	tcs := map[string]struct {
		value1, value2 float64
		expectedResult float64
	}{
		"sub_1": {
			value1: 1, value2: 2,
			expectedResult: 0.5,
		},
		"sub_2": {
			value1: 2, value2: 2,
			expectedResult: 1,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			result := Divide(tc.value1, tc.value2)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
