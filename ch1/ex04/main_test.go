package main

import (
	"testing"
)


func TestCountLine(t *testing.T) {
	var tests = []struct{
		caseName string
		inputTexts []string
		inputFileNames []string
		expectedResults map[string]map[string]int
	}{
		{
			caseName : string("1 file and 1 line"),
			inputTexts : []string{"a"},
			inputFileNames : []string{"sample.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt" : 1}},
		},
		{
			caseName : string("1 file and 2 same line"),
			inputTexts : []string{"a", "a"},
			inputFileNames : []string{"sample.txt", "sample.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt" : 2}},
		},
		{
			caseName : string("1 file and 2 different line"),
			inputTexts : []string{"a", "b"},
			inputFileNames : []string{"sample.txt", "sample.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt" : 1}, "b" : {"sample.txt" : 1}},
		},
		{
			caseName : string("2 file and same line"),
			inputTexts : []string{"a", "a"},
			inputFileNames : []string{"sample.txt", "sample2.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt" : 1, "sample2.txt" : 1}},
		},
		{
			caseName : string("2 file and different line"),
			inputTexts : []string{"a", "b"},
			inputFileNames : []string{"sample.txt", "sample2.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt" : 1}, "b" : {"sample2.txt" : 1}},
		},
		{
			caseName : string("combine 2 files same and different"),
			inputTexts : []string{"a", "b", "b", "c"},
			inputFileNames : []string{"sample.txt", "sample.txt", "sample2.txt", "sample2.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt": 1}, "b" : {"sample.txt" : 1, "sample2.txt" : 1}, "c" : {"sample2.txt" : 1}},
		},
		{
			caseName : string("combine 3 files same and different"),
			inputTexts : []string{"a", "b", "b", "c", "a", "b", "b"},
			inputFileNames : []string{"sample.txt", "sample.txt", "sample2.txt", "sample2.txt", "sample.txt", "sample3.txt", "sample3.txt"},
			expectedResults : map[string]map[string]int{"a" : {"sample.txt": 2}, "b" : {"sample.txt" : 1, "sample2.txt" : 1, "sample3.txt" : 2}, "c" : {"sample2.txt" : 1}},
		},
	}

	for _, testCase := range tests {
		resultMap := make(map[string]map[string]int)
		for i, inputLine := range testCase.inputTexts {
			countLine(inputLine, testCase.inputFileNames[i], resultMap)
		}
		for line, files := range resultMap {
			for file, count := range files {
				if testCase.expectedResults[line][file] != count {
					t.Errorf("error in case %s input value %s and file name %s expects %d but result is %d",
						testCase.caseName, line, file, testCase.expectedResults[line][file], count)
				}
			}
		}
		for line, files := range testCase.expectedResults {
			for file, count := range files {
				if resultMap[line][file] != count {
					t.Errorf("error in case %s input value %s and file name %s expects %d but result is %d",
						testCase.caseName, line, file, testCase.expectedResults[line][file], count)
				}
			}
		}
	}

}

