package main

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Name string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct {
}

type EndList struct {
}

type Decoder struct {
	lex *lexer
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{
		lex: &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}},
	}
	d.lex.scan.Init(r)
	d.lex.next() // get the first token
	return d
}

func (d *Decoder) Token() (Token, error) {
	var t Token
	var err error
	switch d.lex.token {
	case scanner.Ident:
		name := d.lex.text()
		t = Symbol{name}
	case scanner.String:
		s, e := strconv.Unquote(d.lex.text()) // NOTE: ignoring errors
		t = String{s}
		err = e
	case scanner.Int:
		i, e := strconv.Atoi(d.lex.text()) // NOTE: ignoring errors
		t = Int{i}
		err = e
	case '(':
		t = StartList{}
	case ')':
		t = EndList{}
	default:
		panic(fmt.Sprintf("unexpected token %d", d.lex.token))
	}
	d.lex.next()
	return t, err
}
