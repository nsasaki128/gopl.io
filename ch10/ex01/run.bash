#!/usr/bin/env bash
go build ../../ch03/mandelbrot
go build ../../ch01/lissajous
go build .

#png -> jpg, png, gif
./mandelbrot | ./ex01 > mandelbrot_.jpg
./mandelbrot | ./ex01 -e jpg > mandelbrot.jpg
./mandelbrot | ./ex01 -e png > mandelbrot.png
./mandelbrot | ./ex01 -e gif > mandelbrot.gif

#gif -> jpg, png, gif
./lissajous | ./ex01 > lissajous_.jpg
./lissajous | ./ex01 -e jpg > lissajous.jpg
./lissajous | ./ex01 -e png > lissajous.png
./lissajous | ./ex01 -e gif > lissajous.gif


if [ "`sips -g format mandelbrot_.jpg | grep format | cut -f 4 -d \" \"`" != jpeg ]; then
    echo "png to jpg without flag fails"
fi
if [ "`sips -g format mandelbrot.jpg | grep format | cut -f 4 -d \" \"`" != jpeg ]; then
    echo "png to jpg with flag fails"
fi
if [ "`sips -g format mandelbrot.png | grep format | cut -f 4 -d \" \"`" != png ]; then
    echo "png to png fails"
fi
if [ "`sips -g format mandelbrot.gif | grep format | cut -f 4 -d \" \"`" != gif ]; then
    echo "png to gif fails"
fi

if [ "`sips -g format lissajous_.jpg | grep format | cut -f 4 -d \" \"`" != jpeg ]; then
    echo "gif to jpg without flag fails"
fi
if [ "`sips -g format lissajous.jpg | grep format | cut -f 4 -d \" \"`" != jpeg ]; then
    echo "gif to jpg with flag fails"
fi
if [ "`sips -g format lissajous.png | grep format | cut -f 4 -d \" \"`" != png ]; then
    echo "gif to png fails"
fi
if [ "`sips -g format lissajous.gif | grep format | cut -f 4 -d \" \"`" != gif ]; then
    echo "gif to gif fails"
fi

rm -rf ./mandelbrot*
rm -rf ./lissajous*
rm -rf ./ex01

