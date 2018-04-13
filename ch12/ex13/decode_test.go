package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
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
				t.Errorf("Unmarshal %#v want %#v err=nil; but actual %#v err=%#v", test.input, test.want, got, err)
			}
		})
	}
}
