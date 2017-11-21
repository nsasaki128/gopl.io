package weightconv

import "fmt"

type Kilogram float64
type Pound float64


func KToP(k Kilogram) Pound {return Pound(2.20462*k)}
func PToK(p Pound) Kilogram {return Kilogram(p/2.20462)}

//for checking display digit test easier
func (p Pound) String() string { return fmt.Sprintf("%.2flbs", p) }
func (k Kilogram) String() string  { return fmt.Sprintf("%.2fkg", k) }