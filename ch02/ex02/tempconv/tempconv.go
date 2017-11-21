package tempconv
// tempconvパッケージは摂氏（Celsius）と華氏（Fahrenheit）の温度計算を行います。

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)


//for checking display digit test easier
func (c Celsius) String() string { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
