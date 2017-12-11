package main

import (
	"testing"
	"math"
	"net/http/httptest"
	"net/http"
	"fmt"
	"time"
	"math/rand"
)


func TestCreateFlactalInfo(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected flactalInfo
	}{
		{name: "no query", query: "", expected: flactalInfo{x:2.0, y:2.0, m:1.0, f:mandelbrot}},
		{name: "only x", query: "x=3.0", expected: flactalInfo{x:3.0, y:2.0, m:1.0, f:mandelbrot}},
		{name: "only y", query: "y=4.0", expected: flactalInfo{x:2.0, y:4.0, m:1.0, f:mandelbrot}},
		{name: "only m", query: "m=5.0", expected: flactalInfo{x:2.0, y:2.0, m:5.0, f:mandelbrot}},
		{name: "only f", query: "f=newton", expected: flactalInfo{x:2.0, y:2.0, m:1.0, f:newton}},
		{name: "nonexist f", query: "f=neon", expected: flactalInfo{x:2.0, y:2.0, m:1.0, f:mandelbrot}},
		{name: "all", query: "x=3.0&y=4.0&m=5.0&f=newton", expected: flactalInfo{x:3.0, y:4.0, m:5.0, f:newton}},
	}
	rand.Seed(time.Now().UTC().UnixNano())

	testComplex := []complex128{complex(-1.0, -1.0), complex(-1.0, 1.0), complex(1.0, -1.0), complex(1.0, 1.0), complex(0, 0)}
	for i:=0; i < 12; i++{
		testComplex = append(testComplex, complex(rand.Float64()*2.0 , rand.Float64())*1.0)
	}

	//for float error
	const eps= 1e-10
	for _, testCase := range testCases {
		ts := httptest.NewServer(http.HandlerFunc(handler))

		res, _ := http.Get(fmt.Sprintf("%v?%v", ts.URL, testCase.query))
		res.Request.ParseForm()
		actual := createFlactalInfo(res.Request.Form)

		if math.Abs(actual.x-testCase.expected.x) > eps {
			t.Errorf("error in case %s x\nexpected:\t%f\nactual:\t%f\n", testCase.name, testCase.expected.x, actual.x)
			continue
		}
		if math.Abs(actual.y-testCase.expected.y) > eps {
			t.Errorf("error in case %s y\nexpected:\t%f\nactual:\t%f\n", testCase.name, testCase.expected.y, actual.y)
			continue
		}
		if math.Abs(actual.m-testCase.expected.m) > eps {
			t.Errorf("error in case %s m\nexpected:\t%f\nactual:\t%f\n", testCase.name, testCase.expected.m, actual.m)
			continue
		}
		for _, randTest := range testComplex {
			actualR, actualG, actualB, actualA := actual.f(randTest).RGBA()
			expectedR, expectedG, expectedB, expectedA := testCase.expected.f(randTest).RGBA()

			if actualR != expectedR {
				t.Errorf("error in case %s R f\nrand value is %v\nexpected:\t%v\nactual:\t%v\n", testCase.name, randTest, expectedR, actualR)
			}
			if actualG != expectedG{
				t.Errorf("error in case %s G f\nrand value is %v\nexpected:\t%v\nactual:\t%v\n", testCase.name, randTest, expectedG, actualG)
			}
			if actualB != expectedB {
				t.Errorf("error in case %s B f\nrand value is %v\nexpected:\t%v\nactual:\t%v\n", testCase.name, randTest, expectedB, actualB)
			}
			if actualA != expectedA {
				t.Errorf("error in case %s A f\nrand value is %v\nexpected:\t%v\nactual:\t%v\n", testCase.name, randTest,  expectedA, actualA)
			}
		}

	}
}

