package main

import (
	"testing"
	"bytes"
)


func TestWriteDupLineAndFiles(t *testing.T) {
	args := [][]string{{"sample.txt"}}
	results := []string{"2\taaa aaa\tsample.txt\n4\tb\tsample.txt\n"}

	out = new(bytes.Buffer)
	for i, arg := range args{
		writeDupLineAndFiles(arg)
		got := out.(*bytes.Buffer).String()
		result := results[i]
		if got != result {
			t.Errorf("\nresult \t %q\n want \t %s", got, result)
		}
	}

}
