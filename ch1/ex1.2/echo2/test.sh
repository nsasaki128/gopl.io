#!bin/bash
# test contents
test1="test1 empty "
test2="test2 one input "
test3="test3 three inputs "
test4="test4 many inputs "
test5="test5 empty with many white space "
test6="test6 many inputs with many white space "

#test if error happens then show message
if [ "`go run main.go`" != "" ]; then
    echo $test1 "fail"
fi

if [ "`go run main.go hoge`" != "1 hoge" ]; then
    echo $test2 "fail"
fi

if [ "`go run main.go hoge piyo`" != '1 hoge
2 piyo' ]; then
    echo $test3 "fail"
fi

if [ "`go run main.go hoge piyo fuga hogehoge piyopiyo`" != '1 hoge
2 piyo
3 fuga
4 hogehoge
5 piyopiyo' ]; then
    echo $test4 "fail"
fi

if [ "`go run main.go   `" != "" ]; then
    echo $test5 "fail"
fi


if [ "`go run main.go    hoge piyo  fuga   hogehoge  piyopiyo`" != '1 hoge
2 piyo
3 fuga
4 hogehoge
5 piyopiyo' ]; then
    echo $test6 "fail"
fi

echo Done
