package main

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		name string
		isSame bool
	}{
		{name:"same small", isSame:true},
		{name:"different small", isSame:false},
	}

	for _, testCase := range testCases {
		input := [size]int{}
		expected := [size]int{}

		for i:=0; i < size; i++ {
			if testCase.isSame {
				input[i] = 1
				expected[i] = 1
			}else{
				input[i] = i + 1
				expected[i] = size - i
			}
		}
		reverse(&input)
		for i := 0; i < size; i++ {
			if input[i] != expected[i] {
				t.Errorf("error in case %s \nexpected:\t%d\nactual:\t%d\n", testCase.name, expected[i], input[i])
			}
		}

	}
}