package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B){
	inputs := []string{"hoge", "fuga", "piyo", "hogehoge", "fugafuga", "piyopiyo"}

	for i := 0; i < b.N; i++ {
		echo1(inputs)
	}
}

func BenchmarkEcho2(b *testing.B){
	inputs := []string{"hoge", "fuga", "piyo", "hogehoge", "fugafuga", "piyopiyo"}
	for i := 0; i < b.N; i++ {
		echo2(inputs)
	}
}

func BenchmarkEcho3(b *testing.B){
	inputs := []string{"hoge", "fuga", "piyo", "hogehoge", "fugafuga", "piyopiyo"}
	for i := 0; i < b.N; i++ {
		echo3(inputs)
	}
}