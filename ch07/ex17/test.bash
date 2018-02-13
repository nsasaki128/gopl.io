#!bin/bash
# test contents
test1="test1 id: id exist "
test2="test2 id: id does not exist but class exist "
test3="test3 id: id does not exist and class does not exist "
test4="test4 class: class exist "
test5="test5 class: class does not exist but id exist "
test6="test6 class: class does not exist and id does not exist "
test7="test7 id&class: both id and class exist"

go build .

#test if error happens then show message
if [ "`cat test.xml | ./ex17 -i aut`" != "catalog book author: author id sample" ]; then
    echo $test1 "fail"
fi

if [ "`cat test.xml | ./ex17 -i ttl`" != "" ]; then
    echo $test2 "fail"
fi

if [ "`cat test.xml | ./ex17 -i hoge`" != "" ]; then
    echo $test3 "fail"
fi

if [ "`cat test.xml | ./ex17 -c ttl`" != "catalog book title: title class sample" ]; then
    echo $test4 "fail"
fi

if [ "`cat test.xml | ./ex17 -c aut`" != "" ]; then
    echo $test5 "fail"
fi

if [ "`cat test.xml | ./ex17 -c hoge`" != "" ]; then
    echo $test6 "fail"
fi

if [ "`cat test.xml | ./ex17 -i des -c cool`" != "catalog book description: description id and class sample" ]; then
    echo $test7 "fail"
fi
rm -f ./ex17
echo Done
