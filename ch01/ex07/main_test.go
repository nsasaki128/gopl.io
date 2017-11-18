package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

func TestFetch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}))
	defer ts.Close()
	result := new(bytes.Buffer)
	fetch(result, ioutil.Discard, ts.URL)
	if actual, expected := result.String(), "Hello, world!\n"; actual != expected {
		t.Errorf("actual %s; expected %s", actual, expected)
	}

}

/*
func TestFetchError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(nil)
	}))
	defer ts.Close()
	result := new(bytes.Buffer)
	fetch(result, ioutil.Discard, ts.URL)

}
*/