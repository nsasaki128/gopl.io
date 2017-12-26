#!/usr/bin/env bash

URLs=(https://golang.org http://www.gopl.io/ https://github.com/nsasaki128)


go build ../../ch01/fetch/
go build ../findlinks1/
go build .

for URL in ${URLs[@]}; do
    ./fetch $URL | ./findlinks1 > findlinks1Result.txt
    ./fetch $URL | ./ex01 > ex01Result.txt

    diff findlinks1Result.txt ex01Result.txt

    if [ $? == 0 ]; then
        echo $URL" Test Success!"
    else
        echo $URL" Test Fail!"
    fi
done

rm -f ex01 fetch findlinks1 findlinks1Result.txt ex01Result.txt