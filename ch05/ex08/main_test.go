package main

import (
	"testing"

	"strings"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		id            string
		exist         bool
		expectedKey   string
		expectedValue string
	}{
		{name: "no id", input: "<html><head></head><div>test</div><body></body></html>", id: "hoge", exist: false, expectedKey: "", expectedValue: ""},
		{name: "id found", input: `<html><head></head><div id="hoge" sample="fuga">test</div><body></body></html>`, id: "hoge", exist: true, expectedKey: "sample", expectedValue: "fuga"},
		{name: "id not found", input: `<html><head></head><div id="hoge" sample="fuga">test</div><body></body></html>`, id: "fuga", exist: false, expectedKey: "", expectedValue: ""},
		{name: "id not found", input: `<html><head></head><div id="hoge" sample="fuga">test</div><body></body></html>`, id: "sample", exist: false, expectedKey: "", expectedValue: ""},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			doc, _ := html.Parse(strings.NewReader(testCase.input))
			elem := ElementById(doc, testCase.id)

			if elem == nil && testCase.exist {
				t.Errorf("%s expects %s exists but actual does not \n", testCase.input, testCase.id)
			}
			if elem != nil && !testCase.exist {
				t.Errorf("%s expects %s does not exist but actual does \n", testCase.input, testCase.id)
			}
			if elem != nil {
				found := false
				for _, a := range elem.Attr {
					if a.Key == testCase.expectedKey {
						found = true
						if a.Val != testCase.expectedValue {
							t.Errorf("%s id:%s key:%s val expects %s but actual %s \n", testCase.input, testCase.id, testCase.expectedKey, testCase.expectedValue, a.Val)
						}
					}
				}
				if !found {
					t.Errorf("%s id:%s key:%s expects exisitng but actual not \n", testCase.input, testCase.id, testCase.expectedKey)
				}
			}

		})
	}
}
