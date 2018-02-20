#!/usr/bin/env bash
go build main.go
echo "RESULT" > time.txt
echo "start process"
for i in 1 2 4 8 16
do
    echo "start process " ${i}
    start=$( date +%s )
    export GOMAXPROCS=${i}
    ./main -p 4 > result_${i}.png
    end=$( date +%s )
    echo "PROCS " ${i} " takes " $(expr ${end} - ${start}) " sec" >> time.txt
    echo "finish process " ${i}
done
echo "finish process"
rm -f ./main