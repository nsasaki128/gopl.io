#!/usr/bin/env bash
go build main.go
mkdir -p results
./main > results/mandelbrot.png
./main -s supersampling > results/mandelbrotSupersampling.png
./main -s supersampling -d 1.0 > results/mandelbrotSupersamplingBigDif.png
./main -s supersampling -d 0.25 > results/mandelbrotSupersamplingSmallDif.png
rm -f ./main
