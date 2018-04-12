package main

import "testing"

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}
type dataNoLabel struct {
	Labels     []string
	MaxResults int
	Exact      bool
}

func TestPackWithLabel(t *testing.T) {

	tests := []struct {
		name  string
		input data
		want  string
	}{
		{name: "no array", input: data{}, want: "max=0&x=false"},
		{name: "one array", input: data{Labels: []string{"golang"}, MaxResults: 5, Exact: true}, want: "l=golang&max=5&x=true"},
		{name: "some array", input: data{Labels: []string{"golang", "programming", "nsasaki"}, MaxResults: 5, Exact: true}, want: "l=golang&l=programming&l=nsasaki&max=5&x=true"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Pack(&test.input)
			if test.want != got {
				t.Errorf("input %#v want %s got %s\n", test.input, test.want, got)
			}
		})
	}
}
func TestPackNoLabel(t *testing.T) {

	tests := []struct {
		name  string
		input dataNoLabel
		want  string
	}{
		{name: "no array", input: dataNoLabel{}, want: "maxresults=0&exact=false"},
		{name: "one array", input: dataNoLabel{Labels: []string{"golang"}, MaxResults: 5, Exact: true}, want: "labels=golang&maxresults=5&exact=true"},
		{name: "some array", input: dataNoLabel{Labels: []string{"golang", "programming", "nsasaki"}, MaxResults: 5, Exact: true}, want: "labels=golang&labels=programming&labels=nsasaki&maxresults=5&exact=true"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Pack(&test.input)
			if test.want != got {
				t.Errorf("input %#v want %s got %s\n", test.input, test.want, got)
			}
		})
	}
}
