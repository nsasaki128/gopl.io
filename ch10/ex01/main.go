package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var encoder = flag.String("e", "jpg", "encode image type; jpg/png/gif")

func main() {
	flag.Parse()
	img, err := decodeImage(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}

	switch *encoder {
	case "jpg":
		toJPEG(img, os.Stdout)
	case "png":
		toPNG(img, os.Stdout)
	case "gif":
		toGIF(img, os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "Unknown image type <%s>\n", *encoder)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Output format = ", *encoder)
}

func decodeImage(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return img, nil
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 256})
}
