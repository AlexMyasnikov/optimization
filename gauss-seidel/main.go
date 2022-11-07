package main

import (
	"fmt"
	"math"
)

type D struct {
	x1 float64
	x2 float64
}

func (d D) f() float64 {
	//return 2*math.Pow(d.x1, 2) + d.x1*d.x2 + math.Pow(d.x2, 2)
	return 10*math.Pow(d.x1, 2) + 10*math.Pow(d.x2, 2) - 10*d.x1*d.x2 + d.x2
	//return 10 / x1 * (4 - math.Pow(x1, 2))
}

func main() {
	d := D{
		x1: -1,
		x2: 0,
	}
	fmt.Println(d.f())
	// d2 := D{
	// 	x1: 0,
	// 	x2: 0,
	// }
	// d3 := D{
	// 	x1: 0,
	// 	x2: -1,
	// }
	// for {

	// }
}
