package main

import (
	"log"
	"net/http"
	"io"
	"image"
	"math"
	"image/color"
	"strconv"
	"net/url"
	"math/cmplx"
	"image/png"
)

type flactalInfo struct {
	x float64
	y float64
	m float64
	f func(z complex128) color.Color
}


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err!=nil {
		log.Print(err)
	}
	render(w, createFlactalInfo(r.Form))
}

func createFlactalInfo(form url.Values) flactalInfo {
	x  := 2.0
	y  := 2.0
	m  := 1.0
	var f = func(z complex128) color.Color{return mandelbrot(z)}

	for k, v := range form {

		if f64, err := strconv.ParseFloat(v[0], 64); err == nil {
			switch k {
			case "x": x = f64
			case "y": y = f64
			case "m": m = f64
			}
		}
		if k == "f" {
			switch v[0] {
			case "mandelbrot":
				f = func(z complex128) color.Color { return mandelbrot(z) }
			case "newton":
				f = func(z complex128) color.Color { return newton(z) }
			}
		}

	}
	return flactalInfo{x:x, y:y, m:m, f:f}
}

func render(out io.Writer, info flactalInfo) {
	xmin, ymin, xmax, ymax := -math.Abs(info.x/info.m), -math.Abs(info.y/info.m), math.Abs(info.x/info.m), math.Abs(info.y/info.m)
	const(
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, info.f(z))
		}
	}

	png.Encode(out, img) //注意: エンコードエラーを無視
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