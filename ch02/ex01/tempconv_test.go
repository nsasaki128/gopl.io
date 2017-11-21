package tempconv

import (
	"testing"
	"math"
)

var testCases = []struct {
	name string
	k Kelvin
	c Celsius
	f Fahrenheit
}{
	{name:"0K", k:0, c:-273.15, f:-459.67},
	{name:"1K", k:1, c:-272.15, f:-457.87},
	{name:"0°C", k:273.15, c:0, f:32},
}
//for float error
const eps = 1e-10


func TestKToC(t *testing.T) {
	for _, testCase := range testCases {
		if actual := KToC(testCase.k); math.Abs(float64(actual - testCase.c)) > eps{
			t.Errorf("case %s in KToC(%f) expected %f actual %f", testCase.name, testCase.k, testCase.c, actual)
		}
	}
}

func TestKToF(t *testing.T) {
	for _, testCase := range testCases {
		if actual := KToF(testCase.k); math.Abs(float64(actual - testCase.f)) > eps{
			t.Errorf("case %s in KToF(%f) expected %f actual %f", testCase.name, testCase.k, testCase.f, actual)
		}
	}
}

func TestCToF(t *testing.T) {
	for _, testCase := range testCases {
		if actual := CToF(testCase.c); math.Abs(float64(actual - testCase.f)) > eps{
			t.Errorf("case %s in CToF(%f) expected %f actual %f", testCase.name, testCase.c, testCase.f, actual)
		}
	}
}

func TestCToK(t *testing.T) {
	for _, testCase := range testCases {
		if actual := CToK(testCase.c); math.Abs(float64(actual - testCase.k)) > eps{
			t.Errorf("case %s in CToK(%f) expected %f actual %f", testCase.name, testCase.c, testCase.k, actual)
		}
	}
}

func TestFToC(t *testing.T) {
	for _, testCase := range testCases {
		if actual := FToC(testCase.f); math.Abs(float64(actual - testCase.c)) > eps{
			t.Errorf("case %s in FToC(%f) expected %f actual %f", testCase.name, testCase.f, testCase.c, actual)
		}
	}
}

func TestFToK(t *testing.T) {
	for _, testCase := range testCases {
		if actual := FToK(testCase.f); math.Abs(float64(actual - testCase.k)) > eps{
			t.Errorf("case %s in FToK(%f) expected %f actual %f", testCase.name, testCase.f, testCase.k, actual)
		}
	}
}

