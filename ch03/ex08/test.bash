#!/usr/bin/env bash
go build main.go
mkdir -p results
./main -d complex64 > results/mandelbrotComplex64.png
./main -d complex128 > results/mandelbrotComplex128.png
./main -d Float > results/mandelbrotFloat.png
./main -d Rat > results/mandelbrotRat.png
rm -f ./main
