#!/usr/bin/env bash
go build main.go
mkdir -p results
./main > results/sinc.svg
./main -f egg > results/egg.svg
./main -f saddle > results/saddle.svg
rm -f ./main
