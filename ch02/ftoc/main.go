//ftocは華氏（Fahrenheit）から摂氏（Celsius）への変換を二つ表示します。
package main

import "fmt"

const boilingF = 212.0

func main()  {
	const fleezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", fleezingF, fToC(fleezingF)) //32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF)) //212°F = 100°C
	// 出力:
	// boiling point = 212°F or 100°C
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
