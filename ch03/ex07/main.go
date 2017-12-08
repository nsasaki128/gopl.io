// mandelbrot はマンデルブロフラクタルのPNG画像を生成します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
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
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}
// f = z^4 - 1
// f' = 4z^3
// if f(z_i) < eps -> ok
// z_i+1 = z_i - (f(z)/f'(z))
func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	const eps = 1e-10

	var f complex128
	var df complex128
	for n := uint8(0); n < iterations; n++ {
		if cmplx.Abs(complex(1, 0) - z ) < eps {
			return color.RGBA{255-contrast*n, 255, 0, 255}
		}
		if cmplx.Abs(complex(-1, 0) - z ) < eps {
				return color.RGBA{0, 255-contrast*n, contrast*n, 255}
		}
		if cmplx.Abs(complex(0, -1) - z ) < eps {
			return color.RGBA{contrast*n, 0, 255, 255}
		}
		if cmplx.Abs(complex(0, 1) - z ) < eps {
			return color.RGBA{255, contrast*n, 255-contrast*n, 255}
		}
		f = z*z*z*z - 1
		df = 4*z*z*z
		z = z - (f/df)
	}
	return color.Black
}


