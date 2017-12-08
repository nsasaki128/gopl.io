// mandelbrot はマンデルブロフラクタルのPNG画像を生成します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math"
)

func main()  {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
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


