#!/usr/bin/env bash
echo "RESULT" > result.txt
echo "start process"
for i in 1 10 100 1000 10000 100000 1000000 10000000
do
    echo "start process "${i}
    go run main.go -num ${i} >> result.txt
    echo "end process "${i}
done
echo "finish process"
