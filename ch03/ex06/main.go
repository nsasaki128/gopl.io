// mandelbrot はマンデルブロフラクタルのPNG画像を生成します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math"
	"flag"
)
const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height = 1024, 1024
)

var sampling = flag.String("s", "", "function for supersampling")
var dif = flag.Float64("d", 0.5, "dif for supersampling")
func main()  {
	flag.Parse()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if *sampling == "supersampling" {
		renderWithSupersampling(img, *dif)
	} else {
		render(img)
	}
	png.Encode(os.Stdout, img)
}

func render(img *image.RGBA){
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(px, py, mandelbrot(z))
		}
	}
}

func renderWithSupersampling(img *image.RGBA, d float64) {
	const sampleNum = 4
	dx := [sampleNum]float64{-d, -d, +d, +d}
	dy := [sampleNum]float64{-d, +d, -d, +d}
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b uint32 = 0, 0, 0
			for i := 0; i < sampleNum; i++ {
				sx := (float64(px)+dx[i])/width*(xmax-xmin) + xmin
				sy := (float64(py)+dy[i])/height*(ymax-ymin) + ymin
				sz := complex(sx, sy)
				sc := mandelbrot(sz)
				sr, sg, sb, _ := sc.RGBA()
				r += sr
				g += sg
				b += sb
			}
			//because golang color scale is Premultiplied Alpha so each value is multiplied by alpha
			// see the reason https://blog.golang.org/go-image-package
			// see original soruce code https://github.com/golang/go/blob/master/src/image/color/color.go
			r = (r >> 8) / sampleNum
			g = (g >> 8) / sampleNum
			b = (b >> 8) / sampleNum
			c := color.RGBA{uint8(r), uint8(g), uint8(b), 255}
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(px, py, c)
		}
	}
}

//See http://courses.cms.caltech.edu/cs11/material/java/donnie/lab3/index.html
//This pattern uses hsv for iteration num
func mandelbrot(z complex128) color.Color {
	const contrast = 15
	const iterationRange = math.MaxUint8 / contrast
	const iterations = iterationRange*6

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n < iterationRange {
				return color.RGBA{255, uint8(contrast*n), 0, 255}
			} else if n < 2*iterationRange {
				return color.RGBA{uint8(255-contrast*(n-iterationRange)), 255, 0, 255}
			} else if n < 3*iterationRange {
				return color.RGBA{0, 255, uint8(contrast*(n-2*iterationRange)), 255}
			} else if n < 4*iterationRange {
				return color.RGBA{0, uint8(255-contrast*(n-3*iterationRange)), 255, 255}
			} else if n < 5*iterationRange {
				return color.RGBA{uint8(contrast*(n-4*iterationRange)),0 , 255, 255}
			} else {
				return color.RGBA{255, 0, uint8(255-contrast*(n-5*iterationRange)), 255}
			}
		}
	}
	return color.Black
}


