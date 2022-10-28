package main

import (
	"fmt"
	"math"
)

type D struct {
	x1 float64
	x2 float64
	h  float64
	e  float64
}

func f(x1, x2 float64) float64 {
	//return 2*math.Pow(x1, 2) + x1*x2 + math.Pow(x2, 2)
	return 10*math.Pow(x1, 2) + 10*math.Pow(x2, 2) - 10*x1*x2 + x2
	//return 10 / x1 * (4 - math.Pow(x1, 2))
}

func main() {
	d := D{
		x1: -1,
		x2: -1,
		h:  0.2,
		e:  0.01,
	}
	var f0 float64
	var i int32
	var x1, x2 float64
	for {
		fmt.Printf("Iteration: %d\n", i)
		fmt.Printf("x1: %.5f, x2: %.5f\n", d.x1, d.x2)
		fmt.Printf("h: %.5f, e: %.5f\n\n", d.h, d.e)
		i++

		f0 = f(d.x1, d.x2)
		if f(d.x1+d.h, d.x2) < f0 {
			x1 = d.x1 + d.h
		} else if f(d.x1-d.h, d.x2) < f0 {
			x1 = d.x1 - d.h
		} else {
			x1 = d.x1
		}
		f0 = f(x1, d.x2)
		if f(x1, d.x2+d.h) < f0 {
			x2 = d.x2 + d.h
		} else if f(x1, d.x2-d.h) < f0 {
			x2 = d.x2 - d.h
		} else {
			x2 = d.x2
		}
		if d.x1 == x1 && d.x2 == x2 {
			d.h /= 2
		} else {
			for {
				xx1 := d.x1 + 2*(x1-d.x1)
				xx2 := d.x2 + 2*(x2-d.x2)
				if f(xx1, xx2) < f(x1, x2) {
					d.x1, d.x2 = xx1, xx2
					break
				} else {
					d.x1, d.x2 = x1, x2
					break
				}
			}
		}
		if d.h < d.e {
			fmt.Printf("Final: x1 = %.5f, x2 = %.5f, f = %.5f\n", d.x1, d.x2, f(d.x1, d.x2))
			break
		}
	}
}
