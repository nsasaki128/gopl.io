package main
import (
	"fmt"
	. "net"
)
func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDow(v *Flags)      { *v &^= FlagUp }
func SetBroadCast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) //"10001 true
	TurnDow(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) //"10000 false
	SetBroadCast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) //"10010 false
	fmt.Printf("%b %t\n", v, IsCast(v)) //"10010 true
}
