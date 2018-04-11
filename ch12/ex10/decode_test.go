package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	type Result struct {
		Name string
		Comp complex64
		Fl   float32
	}
	var tests = []struct {
		name  string
		input string
		want  interface{}
	}{
		{name: "empty", input: `()`, want: Result{}},
		{name: "include value",
			input: `
(
(Name "hoge") 
(Comp #C(2.0 3.0)) 
(Fl 3.14)
)`, want: Result{Name: "hoge", Comp: complex(2.0, 3.0), Fl: float32(3.14)}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got Result

			dec := NewDecoder(strings.NewReader(test.input))
			err := dec.Decode(&got)
			if err != nil || !reflect.DeepEqual(got, test.want) {
				t.Errorf("dec.Decode %#v want %#v err=nil; but actual %#v err=%#v", test.input, test.want, got, err)
			}
		})
	}
}

func TestUnmarshalBool(t *testing.T) {
	type s struct {
		Flag bool
	}

	tests := []struct {
		name  string
		input string
		want  s
	}{
		{name: "true", input: "((Flag t))", want: s{true}},
		{name: "false", input: "((Flag nil))", want: s{false}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got s
			err := Unmarshal([]byte(test.input), &got)
			if err != nil || !reflect.DeepEqual(test.want, got) {
				t.Errorf("Unmarshal(%s) want %#v, nil and got %#v %#v",
					test.input, test.want, got, err)
			}
		})
	}
}
func TestUnmarshalComplex(t *testing.T) {
	type s struct {
		Val complex64
	}

	tests := []struct {
		name  string
		input string
		want  s
	}{
		{name: "only real", input: "((Val #C(2.0 0)))", want: s{complex(2.0, 0)}},
		{name: "only imag", input: "((Val #C(0 2.0)))", want: s{complex(0, 2.0)}},
		{name: "both real and img", input: "((Val #C(1.0 2.0)))", want: s{complex(1.0, 2.0)}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got s
			err := Unmarshal([]byte(test.input), &got)
			if err != nil || !reflect.DeepEqual(test.want, got) {
				t.Errorf("Unmarshal(%s) want %#v, nil and got %#v %#v",
					test.input, test.want, got, err)
			}
		})
	}
}
func TestUnmarshalInterface(t *testing.T) {
	type s struct {
		Val interface{}
	}

	tests := []struct {
		name  string
		input string
		want  s
	}{
		{name: "array", input: `((Val ("[]int" (1 2 3))))`, want: s{[]int{1, 2, 3}}},
		{name: "slice", input: `((Val ("[3]int" (1 2 3))))`, want: s{[3]int{1, 2, 3}}},
		{name: "map", input: `((Val ("map[int]int" ((1 2)))))`, want: s{map[int]int{1: 2}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got s
			err := Unmarshal([]byte(test.input), &got)
			if err != nil || !reflect.DeepEqual(test.want, got) {
				t.Errorf("Unmarshal(%s) want %#v, nil and got %#v %#v",
					test.input, test.want, got, err)
			}
		})
	}
}
