package main

import (
	"log"
	"net/http"
	"io"
	"math"
	"strconv"
	"net/url"
	"fmt"
	"os"
)

type svgInfo struct {
	width int
	height int
	cells int
	xyrange float64
	color string
	surface string

	xyscale float64
	zscale	float64
}

const (
	angle         = math.Pi / 6         // x, y軸の角度 （=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30度), cos(30度)

var f func(x, y float64) float64

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err!=nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, createSvgInfo(r.Form))
}
func setFunc(info svgInfo){
	switch info.surface {
	case "sinc":
		f = func(x, y float64) float64 {
			r := math.Hypot(x, y) // (0,0)からの距離
			return math.Sin(r) / r
		}
	case "saddle":
		f = func(x, y float64) float64 {
			//x*x-y*y
			scale := 100
			//fix rotation
			xx := x*cos30
			yy := y*sin30
			return ( (xx*xx) - (yy*yy) ) / float64(scale)
		}
	case "egg":
		f = func(x, y float64) float64 {
			//cos(x)cos(y)
			scale := 8
			return math.Cos(x)*math.Cos(y)/float64(scale)
		}
		//bump function
		// see http://math.mit.edu/~stevenj/bump-saddle.pdf
	case "mogul":
		f = func(x, y float64) float64 {
			scale := 8
			r := math.Sin(x)*math.Sin(y)
			if(math.Abs(r-1) < 1e-10 ){
				return 0
			}
			return -math.Pow(math.E, -1/(1 - r*r))/float64(scale)
		}
	default:
		os.Exit(1)
	}

}

func createSvgInfo(form url.Values) svgInfo {
	width, height := 600, 320            // キャンパスの大きさ（画素数）
	cells         := 100                 // 格子のマス目の数
	xyrange       := 30.0                // 軸の範囲（-xyrange..+xyrange）
	color         := "white"
	surface       := "sinc"

	for k, v := range form {
		if i, err := strconv.Atoi(v[0]); err == nil {
			switch k {
			case "width" : width = i
			case "height" : height = i
			case "cells" : cells = i
			}
		}

		if f64, err := strconv.ParseFloat(v[0], 64); err == nil {
			switch k {
			case "xyrange" : xyrange = f64
			}
		}
		switch k {
		case "color" : color = v[0]
		case "surface" :
			switch v[0] {
			case "sinc" : surface = v[0]
			case "egg" : surface = v[0]
			case "saddle" : surface = v[0]
			case "mogul" : surface = v[0]
			}
		}
	}
	xyscale       := float64(width) / 2 / xyrange // x単位 および y単位当たりの画素数
	zscale        := float64(height) * 0.4        // z単位当たりの画素数

	return svgInfo{width:width, height:height, cells:cells, color:color, surface:surface,
	xyrange:xyrange, xyscale:xyscale, zscale:zscale}
}

func surface(out io.Writer, info svgInfo) {
	setFunc(info)
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", info.width, info.height)
	for i := 0; i < info.cells; i++ {
		for j := 0; j < info.cells; j++ {
			ax, ay, ao := corner(i+1, j, info)
			bx, by, bo := corner(i, j, info)
			cx, cy, co := corner(i, j+1, info)
			dx, dy, do := corner(i+1, j+1, info)
			if ao && bo && co && do {
				fmt.Fprintf(out,"<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, info.color)
			}
		}
	}
	fmt.Fprintln(out,"</svg>")
}

func corner(i, j int, info svgInfo) (float64, float64, bool) {
	// マス目(i,j)のかどの点(x,y)を見つける。
	x :=info.xyrange * (float64(i)/float64(info.cells) - 0.5)
	y :=info.xyrange * (float64(j)/float64(info.cells) - 0.5)

	//面の高さzを計算する。
	z := f(x, y)

	// (x,y,z)を2-D SVG キャンパス (sx, sy)へ等角的に投影。
	sx := float64(info.width)/2 + (x-y)*cos30*info.xyscale
	sy := float64(info.height)/2 + (x+y)*sin30*info.xyscale - z*info.zscale

	isFinate := true
	if isInvalid(sx) || isInvalid(sy) {
		isFinate = false
	}

	return sx, sy, isFinate
}


func isInvalid(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}