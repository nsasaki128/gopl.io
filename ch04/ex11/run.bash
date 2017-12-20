#!/usr/bin/env bash
TOKEN=`cat ../../../token.info`

go build main.go
./main -command search
./main -command create -token $TOKEN -title "sample title" -body "sample body"
./main -command update -token $TOKEN -title "sample title update" -body "sample body update" -number 1
./main -command close -token $TOKEN -number 1

rm -f main
