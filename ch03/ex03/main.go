// surfaceは3-D面の関数のSVGレンダリングを計算します。
package main

import (
	"fmt"
	"math"
	"flag"
	"os"
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
var function = flag.String("f", "sinc", "function for visualizing(sinc, saddle, egg, mogul)")

var f func(x, y float64) float64
func init(){
	flag.Parse()
	switch *function {
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
		flag.Usage()
		os.Exit(1)
	}
	
}


func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	min, max := computeMinMax()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ao := corner(i+1, j)
			bx, by, bo := corner(i, j)
			cx, cy, co := corner(i, j+1)
			dx, dy, do := corner(i+1, j+1)
			if ao && bo && co && do {
				color := computeColor(i, j, min, max)
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func computeMinMax () (float64, float64){
	min, max := math.MaxFloat64, -math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x :=xyrange * (float64(i)/cells - 0.5)
			y :=xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if math.IsInf(z, 0) || math.IsNaN(z){
				 continue
			}
			min = math.Min(min, z)
			max = math.Max(max, z)
		}
	}
	return min, max
}

func computeColor (i, j int, min, max float64) string{
	x :=xyrange * (float64(i)/cells - 0.5)
	y :=xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	ratio := (max - z) / (max - min)
	//ratio ranges are from 0 to 1
	red := uint8(255*(1 - ratio))
	blue := uint8(255*ratio)

	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func corner(i, j int) (float64, float64, bool) {
	// マス目(i,j)のかどの点(x,y)を見つける。
	x :=xyrange * (float64(i)/cells - 0.5)
	y :=xyrange * (float64(j)/cells - 0.5)

	//面の高さzを計算する。
	z := f(x, y)

	// (x,y,z)を2-D SVG キャンパス (sx, sy)へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	isFinate := true
	if isInvalid(sx) || isInvalid(sy) {
		isFinate = false
	}

	return sx, sy, isFinate
}


func isInvalid(f float64) bool {
	return math.IsNaN(f) || math.IsInf(f, 0)
}
