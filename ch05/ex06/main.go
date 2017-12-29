// surfaceは3-D面の関数のSVGレンダリングを計算します。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンパスの大きさ（画素数）
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲（-xyrange..+xyrange）
	xyscale       = width / 2 / xyrange // x単位 および y単位当たりの画素数
	zscale        = height * 0.4        // z単位当たりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 （=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30度), cos(30度)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (sx, sy float64) {
	// マス目(i,j)のかどの点(x,y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//面の高さzを計算する。
	z := f(x, y)
	// (x,y,z)を2-D SVG キャンパス (sx, sy)へ等角的に投影。
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // (0,0)からの距離
	return math.Sin(r) / r
}

func isInvalid(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}
