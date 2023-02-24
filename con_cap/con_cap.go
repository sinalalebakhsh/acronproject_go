// Source: https://go.dev/play/p/ZyBKX0T7I2-
// 		https://programming-idioms.org/idiom/250/pick-a-random-value-from-a-map/4435/go


package con_cap

import (
	"math/rand"
)


func ChoiceOfComputer(m map[K]V) K {
	k := rand.Intn(len(m))
	for CountryName , _ := range m {
		if k == 0 {
			return CountryName 
		}
		k--
	}
	panic("unreachable")
}

type K string
type V string

var CountriesCapital = map[K]V{
	"AFGHANISTAN": "KABUL",
	"ALBANIA": "TIRANA",
	"ALGERIA": "ALGIERS",
	"ANDORRA": "ANDORRA LA VELLA",
	"ANGOLA": "LUANDA",
	"ANTIGUA & BARBUDA": "SAINT JOHN'S",
	"ARGENTINA": "BUENOS AIRES",
	"ARMENIA": "YEREVAN",
	"AUSTRALIA": "CANBERRA",
	"AUSTRIA": "VIENNA",
	"AZERBAIJAN": "BAKU",
	// -------------------------------- 
	"afghanistan": "kabul",
	"albania": "tirana",
	"algeria": "algiers",
	"andorra": "andorra la vella",
	"angola": "luanda",
	"antigua & barbuda": "saint john's",
	"argentina": "buenos aires",
	"armenia": "yerevan",
	"australia": "canberra",
	"austria": "vienna",
	"azerbaijan": "baku",
}


/*

	// Source: https://go.dev/play/p/ZyBKX0T7I2-
	// 		https://programming-idioms.org/idiom/250/pick-a-random-value-from-a-map/4435/go

	package main

	import (
		"fmt"
		"math/rand"
	)

	type K string
	type V string

	func pick(m map[K]V) V {
		k := rand.Intn(len(m))
		for _, x := range m {
			if k == 0 {
				return x
			}
			k--
		}
		panic("unreachable")
	}

	func main() {
		m := map[K]V{
			"one":   "un",
			"two":   "deux",
			"three": "trois",
			"four":  "quatre",
		}

		x := pick(m)

		fmt.Println(x)
	}
*/


