package main

import (
	"testing"
	"math"
	"net/http/httptest"
	"net/http"
	"fmt"
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

func TestCreateSvgInfo(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected svgInfo
	}{
		{name: "no query", query: "", expected: svgInfo{width:600, height:320, cells:100, xyrange:30.0, color:"white", surface:"sinc"}},
		{name: "only width", query: "width=200", expected: svgInfo{width:200, height:320, cells:100, xyrange:30.0, color:"white", surface:"sinc"}},
		{name: "only height", query: "height=300", expected: svgInfo{width:600, height:300, cells:100, xyrange:30.0, color:"white", surface:"sinc"}},
		{name: "only cells", query: "cells=10", expected: svgInfo{width:600, height:320, cells:10, xyrange:30.0, color:"white", surface:"sinc"}},
		{name: "only xyrange", query: "xyrange=20.0", expected: svgInfo{width:600, height:320, cells:100, xyrange:20.0, color:"white", surface:"sinc"}},
		{name: "only color", query: "color=black", expected: svgInfo{width:600, height:320, cells:100, xyrange:30.0, color:"black", surface:"sinc"}},
		{name: "only surface", query: "surface=egg", expected: svgInfo{width:600, height:320, cells:100, xyrange:30.0, color:"white", surface:"egg"}},
		{name: "miss surface", query: "surface=hoge", expected: svgInfo{width:600, height:320, cells:100, xyrange:30.0, color:"white", surface:"sinc"}},
		{name: "all", query: "width=200&height=300&cells=10&xyrange=20.0&color=black&surface=saddle", expected: svgInfo{width:200, height:300, cells:10, xyrange:20.0, color:"black", surface:"saddle"}},
	}
	//for float error
	const eps= 1e-10
	for _, testCase := range testCases {
		ts := httptest.NewServer(http.HandlerFunc(handler))

		res, _ := http.Get(fmt.Sprintf("%v?%v", ts.URL, testCase.query))
		res.Request.ParseForm()
		actual := createSvgInfo(res.Request.Form)
		if actual.width != testCase.expected.width {
			t.Errorf("error in case %s width\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.width, actual.width)
			continue
		}
		if actual.height != testCase.expected.height {
			t.Errorf("error in case %s height\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.height, actual.height)
			continue
		}
		if actual.cells != testCase.expected.cells {
			t.Errorf("error in case %s cells\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.cells, actual.cells)
			continue
		}
		if math.Abs(actual.xyrange-testCase.expected.xyrange) > eps {
			t.Errorf("error in case %s xyrange\nexpected:\t%f\nactual:\t%f\n", testCase.name, testCase.expected.xyrange, actual.xyrange)
			continue
		}
		if actual.color != testCase.expected.color {
			t.Errorf("error in case %s color\nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected.color, actual.color)
			continue
		}
		if actual.surface != testCase.expected.surface {
			t.Errorf("error in case %s surface\nexpected:\t%s\nactual:\t%s\n", testCase.name, testCase.expected.surface, actual.surface)
			continue
		}
	}
}
