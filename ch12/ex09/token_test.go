package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDecoder_Token(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  []Token
	}{
		{name: "empty", input: `()`, want: []Token{StartList{}, EndList{}}},
		{name: "only Int", input: `(3)`, want: []Token{StartList{}, Int{3}, EndList{}}},
		{name: "only String", input: `("hoge")`, want: []Token{StartList{}, String{"hoge"}, EndList{}}},
		{name: "only Symbol", input: `(hoge)`, want: []Token{StartList{}, Symbol{"hoge"}, EndList{}}},
		{name: "nest values", input: `(hoge (fuga 23))`,
			want: []Token{StartList{}, Symbol{"hoge"}, StartList{}, Symbol{"fuga"}, Int{23}, EndList{}, EndList{}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := NewDecoder(strings.NewReader(test.input))
			for _, want := range test.want {
				got, err := d.Token()
				if !reflect.DeepEqual(got, want) || err != nil {
					t.Fatal(fmt.Errorf("d.Decode %#v want %#v err=nil; but actual %#v err=%#v\n", test.input, want, got, err))
				}
			}
		})
	}
}
