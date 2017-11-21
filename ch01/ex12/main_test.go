package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"math"
)

func TestCreateLissajousInfo(t *testing.T) {
	testCases := []struct {
		name string
		query    string
		expected lissajousInfo
	}{
		{name:"no query", query   : "", expected:lissajousInfo{cycles:5, res:0.001, size:100, nframes:64, delay:8}},
		{name:"only cycles query", query   : "cycles=3", expected:lissajousInfo{cycles:3, res:0.001, size:100, nframes:64, delay:8}},
		{name:"only res query", query   : "res=0.002", expected:lissajousInfo{cycles:5, res:0.002, size:100, nframes:64, delay:8}},
		{name:"only size query",query   : "size=200", expected:lissajousInfo{cycles:5, res:0.001, size:200, nframes:64, delay:8}},
		{name:"only nframes query",query   : "nframes=32", expected:lissajousInfo{cycles:5, res:0.001, size:100, nframes:32, delay:8}},
		{name:"only delay query",query   : "delay=3", expected:lissajousInfo{cycles:5, res:0.001, size:100, nframes:64, delay:3}},
		{name:"full query",query   : "cycles=3&res=0.002&size=200&nframes=32&delay=3", expected:lissajousInfo{cycles:3, res:0.002, size:200, nframes:32, delay:3}},
		{name:"duplicated cycles query", query   : "cycles=3&cycles=20", expected:lissajousInfo{cycles:3, res:0.001, size:100, nframes:64, delay:8}},
		}
	//for float error
	const eps = 1e-10
	for _, testCase := range testCases{
			ts := httptest.NewServer(http.HandlerFunc(handler))
			defer ts.Close()

			res, _ := http.Get(fmt.Sprintf("%v?%v", ts.URL, testCase.query))
			res.Request.ParseForm()
			actual := createLissajousInfo(res.Request.Form)
			if actual.cycles != testCase.expected.cycles {
				t.Errorf("error in case %s cycles\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.cycles, actual.cycles)
				continue
			}
			if math.Abs(actual.res - testCase.expected.res) > eps {
				t.Errorf("error in case %s res\nexpected:\t%f\nactual:\t%f\n", testCase.name, testCase.expected.res, actual.res)
				continue
			}
			if actual.size != testCase.expected.size {
				t.Errorf("error in case %s size\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.size, actual.size)
				continue
			}
			if actual.nframes != testCase.expected.nframes {
				t.Errorf("error in case %s nframes\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.nframes, actual.nframes)
				continue
			}
			if actual.delay != testCase.expected.delay {
				t.Errorf("error in case %s delay\nexpected:\t%d\nactual:\t%d\n", testCase.name, testCase.expected.delay, actual.delay)
				continue
			}
	}





}
