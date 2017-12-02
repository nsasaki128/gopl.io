package main

import (
	"testing"
	"math"
)

func TestIsInvalid(t *testing.T) {
	var testCases = []struct {
		name     string
		in       float64
		expected bool
	}{
		{name: "normal case 0", in: 0,
			expected: true,
		},
		{name: "normal case -1", in: -1,
			expected: true,
		},
		{name: "normal case 1", in: 1,
			expected: true,
		},
		{name: "normal case pi", in: math.Pi,
			expected: true,
		},
		{name: "invaild case const NaN", in: math.NaN(),
			expected: false,
		},
		{name: "invalid case divide 0 by 0", in: func(f float64) float64{return f/f}(0),
			expected: false,
		},
		{name: "infinite positive value", in: math.Inf(1),
			expected: false,
		},
		{name: "infinite negative value", in: math.Inf(-1),
			expected: false,
		},
	}
	for _, testCase := range testCases {
		result := isInvalid(testCase.in)
		if result == testCase.expected {
			t.Errorf("case %s expected %t actual %t", testCase.name, testCase.expected, result)
		}
	}
}

