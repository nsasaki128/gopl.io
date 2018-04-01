// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/package main
package main

import (
	"io"
	"testing"
)

// NOTE: we can't use !+..!- comments to excerpt these tests
// into the book because it defeats the Example mechanism,
// which requires the // Output comment to be at the end
// of the function.

func Example_slice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func Example_struct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func Example_array() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

//Add for ex01
func Example_mapArray() {
	mapArrayTest := map[[2]string]int{[2]string{"00", "01"}: 1, [2]string{"01", "02"}: 2}
	Display("mapArray", mapArrayTest)
	// Unordered output:
	// Display mapArray (map[[2]string]int):
	// mapArray[[2]string{"00", "01"}] = 1
	// mapArray[[2]string{"01", "02"}] = 2
}
func Example_mapStruct() {
	type testStruct struct {
		Id   int
		Name string
	}
	mapStructTest := map[testStruct]int{
		testStruct{Id: 1, Name: "01"}: 2,
		testStruct{Id: 2, Name: "02"}: 3,
	}
	Display("mapStruct", mapStructTest)
	// Unordered output:
	// Display mapStruct (map[main.testStruct]int):
	// mapStruct[main.testStruct{Id:1, Name:"01"}] = 2
	// mapStruct[main.testStruct{Id:2, Name:"02"}] = 3
}

func Example_movie() {
	//!+movie
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	//!-movie
	//!+strangelove
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
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
	//!-strangelove
	Display("strangelove", strangelove)

	// We don't use an Output: comment since displaying
	// a map is nondeterministic.
	/*
		//!+output
		Display strangelove (display.Movie):
		strangelove.Title = "Dr. Strangelove"
		strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
		strangelove.Year = 1964
		strangelove.Color = false
		strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
		strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
		strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
		strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
		strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
		strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
		strangelove.Oscars[0] = "Best Actor (Nomin.)"
		strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
		strangelove.Oscars[2] = "Best Director (Nomin.)"
		strangelove.Oscars[3] = "Best Picture (Nomin.)"
		strangelove.Sequel = nil
		//!-output
	*/
}

// This test ensures that the program terminates without crashing.
func Test(t *testing.T) {
	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*(*c.Tail).Tail).Tail).Value = 42
	// (*(*(*(*c.Tail).Tail).Tail).Tail).Value = 42

}
