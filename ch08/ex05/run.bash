#!/usr/bin/env bash
echo "RESULT" > time.txt
echo "start process"
for i in 2 3 4 5 6 7 8 9
do
    echo "start process " ${i}
    start=$( date +%s )
    go run main.go -p ${i} > result_${i}.png
    end=$( date +%s )
    echo "para " ${i} " takes " $(expr ${end} - ${start}) " sec" >> time.txt
    echo "finish process " ${i}
done
echo "finish process"
