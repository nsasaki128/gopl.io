// mandelbrot はマンデルブロフラクタルのPNG画像を生成します。
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"math/big"
	"flag"
)
const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height = 1024, 1024
)

type ComplexBigFloat struct{
	re big.Float
	im big.Float
}
type ComplexBigRat struct{
	re big.Rat
	im big.Rat
}
var dataType = flag.String("d", "complex64", "function for data type (complex64, complex128, Float, Rat)")
func main()  {

	flag.Parse()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	switch *dataType {
	case "complex64":
		renderComplex64(img)
	case "complex128":
		renderComplex128(img)
	case "Float":
		renderBigFloat(img)
	case "Rat":
		renderBigRat(img)
	default:
		flag.Usage()
		os.Exit(1)
	}
	png.Encode(os.Stdout, img)
}

func renderComplex64(img *image.RGBA){
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
}
func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		re, im := real(v), imag(v)
		if re*re + im*im - 4 > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderComplex128(img *image.RGBA){
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
}

func mandelbrotComplex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderBigFloat(img *image.RGBA){
	var widthF, heightF, xminF, xdistF, yminF, ydistF big.Float
	widthF.SetInt64(width)
	heightF.SetInt64(height)

	xminF.SetInt64(xmin)
	xdistF.SetInt64(xmax - xmin)

	yminF.SetInt64(ymin)
	ydistF.SetInt64(ymax - ymin)


	var y, x big.Float
	for py := int64(0); py < height; py++ {
		//y := float64(py)/height*(ymax-ymin) + ymin
		y.SetInt64(py)
		y.Quo(&y, &heightF)
		y.Mul(&y, &ydistF)
		y.Add(&y, &yminF)

		for px := int64(0); px < width; px++ {
			//x := float64(px)/width*(xmax-xmin) + xmin
			x.SetInt64(px)
			x.Quo(&x, &widthF)
			x.Mul(&x, &xdistF)
			x.Add(&x, &xminF)

			z := ComplexBigFloat{x, y}
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(int(px), int(py), mandelbrotBigFloat(z))
		}
	}
}
func mandelbrotBigFloat(z ComplexBigFloat) color.Color {
	const iterations = 64
	const contrast = 15
	const absVal = 2
	var absValSq big.Float
	absValSq.SetInt64(absVal*absVal)

	var vx, vy big.Float
	for n := uint8(0); n < iterations; n++ {
		var vxx, vyy, vxy, res big.Float
		vxx.Mul(&vx, &vx)
		vyy.Mul(&vy, &vy)
		vxy.Mul(&vx, &vy)

		//real x*x-y*y
		//imagin 2*x*y
		vx.Sub(&vxx, &vyy)
		vy.Add(&vxy, &vxy)

		vx.Add(&vx, &z.re)
		vy.Add(&vy, &z.im)

		vxx.Mul(&vx, &vx)
		vyy.Mul(&vy, &vy)
		res.Add(&vxx, &vyy)
		ret:= res.Cmp(&absValSq)

		if ret > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderBigRat(img *image.RGBA){
	var widthF, heightF, xminF, xdistF, yminF, ydistF big.Rat
	widthF.SetInt64(width)
	heightF.SetInt64(height)

	xminF.SetInt64(xmin)
	xdistF.SetInt64(xmax - xmin)

	yminF.SetInt64(ymin)
	ydistF.SetInt64(ymax - ymin)


	var y, x big.Rat
	for py := int64(0); py < height; py++ {
		//y := float64(py)/height*(ymax-ymin) + ymin
		y.SetInt64(py)
		y.Quo(&y, &heightF)
		y.Mul(&y, &ydistF)
		y.Add(&y, &yminF)

		for px := int64(0); px < width; px++ {
			//x := float64(px)/width*(xmax-xmin) + xmin
			x.SetInt64(px)
			x.Quo(&x, &widthF)
			x.Mul(&x, &xdistF)
			x.Add(&x, &xminF)

			z := ComplexBigRat{x, y}
			//画像の点（px, py）は複素数値 z を表している。
			img.Set(int(px), int(py), mandelbrotBigRat(z))
		}
	}
}
func mandelbrotBigRat(z ComplexBigRat) color.Color {
	const iterations = 8
	const contrast = 15
	const absVal = 2
	var absValSq big.Rat
	absValSq.SetInt64(absVal*absVal)

	var vx, vy big.Rat
	for n := uint8(0); n < iterations; n++ {
		var vxx, vyy, vxy, res big.Rat
		vxx.Mul(&vx, &vx)
		vyy.Mul(&vy, &vy)
		vxy.Mul(&vx, &vy)

		//real x*x-y*y
		//imagin 2*x*y
		vx.Sub(&vxx, &vyy)
		vy.Add(&vxy, &vxy)

		vx.Add(&vx, &z.re)
		vy.Add(&vy, &z.im)

		vxx.Mul(&vx, &vx)
		vyy.Mul(&vy, &vy)
		res.Add(&vxx, &vyy)
		ret:= res.Cmp(&absValSq)

		if ret > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
