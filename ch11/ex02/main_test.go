package main

import (
	"reflect"
	"testing"
)

func toIntSet(args ...int) *IntSet {
	s := IntSet{}
	for _, arg := range args {
		s.Add(arg)
	}
	return &s
}

func (s *IntSet) eq(t *IntSet) bool {
	for i, _ := range s.words {
		if len(t.words) <= i {
			if s.words[i] != 0 {
				return false
			}
		} else {
			if s.words[i] != t.words[i] {
				return false
			}
		}
	}
	for i, _ := range t.words {
		if len(s.words) <= i {
			if t.words[i] != 0 {
				return false
			}
		}
	}
	return true
}

// This test is used for ch11 ex02
func TestAdd(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{name: "nothing", input: nil},
		{name: "1 input", input: []int{1}},
		{name: "several inputs", input: []int{0, 1, 2, 3, 4, 5, 123, 144, 255, 256}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var x IntSet
			y := make(map[int]bool)
			for _, i := range testCase.input {
				x.Add(i)
				y[i] = true

			}
			for i := 0; i < 1024; i++ {
				if x.Has(i) != y[i] {
					t.Errorf("Error in case %s, x.Has(%d)=%t and y[%d]=%t, these are expected as same.", testCase.name, i, x.Has(i), i, y[i])
				}
			}

		})
	}

}

func TestLen(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		expected int
	}{
		{name: "nothing", input: toIntSet([]int{}...), expected: 0},
		{name: "1 input", input: toIntSet([]int{1}...), expected: 1},
		{name: "several inputs", input: toIntSet([]int{0, 1, 2, 3, 4, 5, 123, 144, 255, 256}...), expected: 10},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.input.Len()
			if actual != testCase.expected {
				t.Errorf("input %v expects %d but actual %d\n", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		remove   []int
		expected *IntSet
	}{
		{name: "empty", input: toIntSet([]int{}...), remove: []int{}, expected: toIntSet([]int{}...)},
		{name: "nothing 1 input", input: toIntSet([]int{1}...), remove: []int{}, expected: toIntSet([]int{1}...)},
		{name: "nothing remove", input: toIntSet([]int{1, 63, 64}...), remove: []int{3}, expected: toIntSet([]int{1, 63, 64}...)},
		{name: "remove 1 input", input: toIntSet([]int{1, 63, 64}...), remove: []int{1}, expected: toIntSet([]int{63, 64}...)},
		{name: "remove several input", input: toIntSet([]int{1, 63, 64, 127, 128, 255, 256}...), remove: []int{1, 23, 127, 128, 255, 256}, expected: toIntSet([]int{63, 64}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			for _, r := range testCase.remove {
				testCase.input.Remove(r)
			}
			if !testCase.input.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.input)
			}
		})
	}
}

func TestClear(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		expected *IntSet
	}{
		{name: "empty", input: toIntSet([]int{}...), expected: toIntSet([]int{}...)},
		{name: "1 input", input: toIntSet([]int{1}...), expected: toIntSet([]int{}...)},
		{name: "several input", input: toIntSet([]int{1, 63, 64, 127, 128, 255, 256}...), expected: toIntSet([]int{}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.input.Clear()
			if !testCase.input.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.input)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		expected *IntSet
	}{
		{name: "empty", input: toIntSet([]int{}...), expected: toIntSet([]int{}...)},
		{name: "1 input", input: toIntSet([]int{1}...), expected: toIntSet([]int{1}...)},
		{name: "several input", input: toIntSet([]int{1, 63, 64, 127, 128, 255, 256}...), expected: toIntSet([]int{1, 63, 64, 127, 128, 255, 256}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.input.Copy()
			if !actual.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, actual)
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		addList  []int
		expected *IntSet
	}{
		{name: "empty", input: toIntSet([]int{}...), addList: []int{}, expected: toIntSet([]int{}...)},
		{name: "1 input", input: toIntSet([]int{1}...), addList: []int{2}, expected: toIntSet([]int{1, 2}...)},
		{name: "several input", input: toIntSet([]int{1, 63, 64, 127, 128, 255, 256}...), addList: []int{2, 3, 4, 257}, expected: toIntSet([]int{1, 2, 3, 4, 63, 64, 127, 128, 255, 256, 257}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.input.AddAll(testCase.addList...)
			if !testCase.input.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.input)
			}
		})
	}
}

func TestIntersectWith(t *testing.T) {
	testCases := []struct {
		name     string
		inputSrc *IntSet
		inputDst *IntSet
		expected *IntSet
	}{
		{name: "empty", inputSrc: toIntSet([]int{}...), inputDst: toIntSet([]int{}...), expected: toIntSet([]int{}...)},
		{name: "same", inputSrc: toIntSet([]int{1}...), inputDst: toIntSet([]int{1}...), expected: toIntSet([]int{1}...)},
		{name: "1 intersect", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{1, 3}...), expected: toIntSet([]int{1}...)},
		{name: "no intersect", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{3, 4}...), expected: toIntSet([]int{}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.inputSrc.IntersectWith(testCase.inputDst)
			if !testCase.inputSrc.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.inputSrc)
			}
		})
	}
}

func TestDifferenceWith(t *testing.T) {
	testCases := []struct {
		name     string
		inputSrc *IntSet
		inputDst *IntSet
		expected *IntSet
	}{
		{name: "empty", inputSrc: toIntSet([]int{}...), inputDst: toIntSet([]int{}...), expected: toIntSet([]int{}...)},
		{name: "same", inputSrc: toIntSet([]int{1}...), inputDst: toIntSet([]int{1}...), expected: toIntSet([]int{}...)},
		{name: "1 difference", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{1, 3}...), expected: toIntSet([]int{2}...)},
		{name: "no difference", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{3, 4}...), expected: toIntSet([]int{1, 2}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.inputSrc.DifferenceWith(testCase.inputDst)
			if !testCase.inputSrc.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.inputSrc)
			}
		})
	}
}

func TestSynmetricWith(t *testing.T) {
	testCases := []struct {
		name     string
		inputSrc *IntSet
		inputDst *IntSet
		expected *IntSet
	}{
		{name: "empty", inputSrc: toIntSet([]int{}...), inputDst: toIntSet([]int{}...), expected: toIntSet([]int{}...)},
		{name: "same", inputSrc: toIntSet([]int{1}...), inputDst: toIntSet([]int{1}...), expected: toIntSet([]int{}...)},
		{name: "1 difference", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{1, 3}...), expected: toIntSet([]int{2, 3}...)},
		{name: "no difference", inputSrc: toIntSet([]int{1, 2}...), inputDst: toIntSet([]int{3, 4}...), expected: toIntSet([]int{1, 2, 3, 4}...)},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.inputSrc.SynmetricWith(testCase.inputDst)
			if !testCase.inputSrc.eq(testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, testCase.inputSrc)
			}
		})
	}
}

func TestElems(t *testing.T) {
	testCases := []struct {
		name     string
		input    *IntSet
		expected []uint
	}{
		{name: "empty", input: toIntSet([]int{}...), expected: []uint{}},
		{name: "1 input", input: toIntSet([]int{1}...), expected: []uint{1}},
		{name: "many input", input: toIntSet([]int{1, 2, 63, 64, 256}...), expected: []uint{1, 2, 63, 64, 256}},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.input.Elems()
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("expects %v but actual %v\n", testCase.expected, actual)
			}
		})
	}
}
