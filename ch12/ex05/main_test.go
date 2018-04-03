// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"encoding/json"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	var movie Movie
	if err := json.Unmarshal(data, &movie); err != nil {
		t.Fatalf("JSON Unmarshal failed: %v", err)
	}
	t.Logf("json.Unmarshal() = \n%+v\n", movie)

	// Decode it
	/*
		var movie Movie
		if err := Unmarshal(data, &movie); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		t.Logf("Unmarshal() = %+v\n", movie)

		// Check equality.
		if !reflect.DeepEqual(movie, strangelove) {
			t.Fatal("not equal")
		}

		// Pretty-print it:
		data, err = MarshalIndent(strangelove)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("MarshalIdent() = %s\n", data)
	*/
}

var tests = []struct {
	name  string
	input interface{}
	want  string
}{
	{name: "bool true", input: true, want: "true"},
	{name: "bool false", input: false, want: "false"},
	{name: "float positive", input: 1.0, want: "1.000000"},
	{name: "float negative", input: -1.0, want: "-1.000000"},
	{name: "float positive zero", input: 0.0, want: "0.000000"},
	{name: "float negative zero", input: -0.0, want: "0.000000"},
	{name: "map", input: map[string]int{"a": 1, "b": 2}, want: `{"a": 1, "b": 2}`},
	{name: "array", input: struct{ v interface{} }{[]int{1, 2, 3}}, want: `{"v": [1, 2, 3]}`},
	{name: "interface nil", input: struct{ v interface{} }{nil}, want: `{"v": null}`},
}

func TestMarshal(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Marshal(test.input)
			if err != nil || string(got) != test.want {
				t.Errorf("Marshal(%v) = \n%q, %v\n want %q, nil", test.input, got, err, test.want)
			}
		})
	}
}
