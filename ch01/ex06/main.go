package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black}

const (
	variation = 64 //パレットの色の数
)

func addColor(num int) {
	r := 0xff0000
	g := 0x00ff00
	b := 0x0000ff

	max     := 0xffffff
	dif     := max/num

	current := 0x000000

	for i := 0; i < num; i++ {
		current += dif
		palette = append(palette,
			color.RGBA{uint8(current&r>>16), uint8(current&g>>8), uint8(current&b), 0xff},
		)
	}

}

func main() {
	addColor(variation)
	rand.Seed(time.Now().UTC().UnixNano())

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発信器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンパスは [-size..+size] の範囲で扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)

	freq := rand.Float64() * 3.0 // 発信器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img  := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			thisPoint := uint8(((i+int(t)) % (len(palette)-1)) + 1)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), thisPoint)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) //注意: エンコードエラーを無視
}

