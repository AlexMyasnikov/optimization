package main

import (
	"fmt"
	"math"
)

func (d *D) f() float64 {
	//return 2*math.Pow(d.x1, 2) + d.x1*d.x2 + math.Pow(d.x2, 2)
	return 10*math.Pow(d.x1, 2) + 10*math.Pow(d.x2, 2) - 10*d.x1*d.x2 + d.x2
}

type D struct {
	x1 float64
	x2 float64
	h  float64
	e  float64
}

func main() {
	d := D{
		x1: -1,
		x2: -1,
		h:  0.2,
		e:  0.01,
	}
	f0 := d.f()
	var f1 float64
	var i int32
	for {
		i++
		d.x1 += d.h
		f1 = d.f()
		if f1 > f0 {
			d.x1 -= 2 * d.h // возвращаем x1 на место и уменьшаем на h
			f1 = d.f()
			if f1 > f0 {
				d.x1 += d.h
				d.h /= 2
			}
		}
		f0 = f1
		d.x2 += d.h
		f1 = d.f()
		if f1 > f0 {
			d.x2 -= 2 * d.h // возвращаем x2 на место и уменьшаем на h
			f1 = d.f()
			if f1 > f0 {
				d.x2 += d.h
				d.h /= 2
			}
		}
		fmt.Printf("Iteration: %d\n", i)
		fmt.Printf("x1: %.5f, x2: %.5f\n", d.x1, d.x2)
		fmt.Printf("h: %.5f, e: %.5f\n", d.h, d.e)
		fmt.Printf("f0: %.5f, f1: %.5f\n\n", f0, f1)
		if d.h < d.e {
			fmt.Printf("Final: x1 = %.5f, x2 = %.5f, f = %.5f\n", d.x1, d.x2, d.f())
			break
		}
	}
}
