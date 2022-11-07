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

// центр тяжести
func gravity(dBest, dMiddle D) D {
	x1 := (dBest.x1 + dMiddle.x1) / 2
	x2 := (dBest.x2 + dMiddle.x2) / 2
	return D{x1: x1, x2: x2}
}

// if delta <= epsilon --> stop the optimization
func delta(grav, fBest, fMiddle, fWorst float64) float64 {
	var sum float64
	for _, v := range []float64{fBest, fMiddle, fWorst} {
		sum += math.Pow(v-grav, 2)
	}
	return math.Sqrt(sum / float64(2+1))
}

// отражение
func reflection(alpha float64, grav, dWorst D) D {
	x1 := grav.x1 + alpha*(grav.x1-dWorst.x1)
	x2 := grav.x2 + alpha*(grav.x2-dWorst.x2)
	return D{x1: x1, x2: x2}
}

// растяжение
func stretch(gamma float64, grav, refl D) D {
	x1 := grav.x1 + gamma*(refl.x1-grav.x1)
	x2 := grav.x2 + gamma*(refl.x2-grav.x2)
	return D{x1: x1, x2: x2}
}

// сжатие
func compress(beta float64, grav, dWorst D) D {
	x1 := grav.x1 + beta*(dWorst.x1-grav.x1)
	x2 := grav.x2 + beta*(dWorst.x2-grav.x2)
	return D{x1: x1, x2: x2}
}

// редукция
func reduction(beta float64, dBest, d D) D {
	d.x1 = dBest.x1 + beta*(d.x1-dBest.x1)
	d.x2 = dBest.x2 + beta*(d.x2-dBest.x2)
	return d
}

func main() {
	e := 0.0001
	alpha, beta, gamma := 1.0, 0.5, 2.0
	d1 := D{
		x1: -1,
		x2: 0,
	}
	d2 := D{
		x1: 0,
		x2: 0,
	}
	d3 := D{
		x1: 0,
		x2: -1,
	}
	i := 1
	var dBest, dWorst, dMiddle D
	for {
		fmt.Printf("Iteration: %d\n", i)
		i++
		f1 := d1.f()
		f2 := d2.f()
		f3 := d3.f()
		switch fmin := math.Min(f1, math.Min(f2, f3)); fmin {
		case f1:
			dBest = d1
		case f2:
			dBest = d2
		case f3:
			dBest = d3
		}
		switch fmax := math.Max(f1, math.Max(f2, f3)); fmax {
		case f1:
			dWorst = d1
		case f2:
			dWorst = d2
		case f3:
			dWorst = d3
		}
		if d1 != dBest && d1 != dWorst {
			dMiddle = d1
		} else if d2 != dBest && d2 != dWorst {
			dMiddle = d2
		} else {
			dMiddle = d3
		}
		grav := gravity(dBest, dMiddle)
		if delta(grav.f(), dBest.f(), dMiddle.f(), dWorst.f()) <= e {
			fmt.Printf("x1: %.5f, x2: %.5f\n", dBest.x1, dBest.x2)
			fmt.Printf("f(xBest) = %.5f\n\n", dBest.f())
			break
		}
		refl := reflection(alpha, grav, dWorst)
		if refl.f() < dBest.f() {
			fmt.Println("Stretch")
			dNew := stretch(gamma, grav, refl)
			fmt.Println("New simplix")
			if dNew.f() < dBest.f() {
				dWorst = dNew
			} else {
				dWorst = refl
			}
		} else {
			if refl.f() <= dMiddle.f() {
				fmt.Println("New simplix")
				dWorst = refl
			} else {
				if refl.f() <= dWorst.f() {
					fmt.Println("Compress")
					dWorst = compress(beta, grav, dWorst)
				} else {
					fmt.Println("Reduction")
					dWorst = reduction(beta, dBest, dWorst)
					dMiddle = reduction(beta, dBest, dMiddle)
				}
			}
		}
		d1 = dBest
		d2 = dWorst
		d3 = dMiddle
		fmt.Printf("x1: %.5f, x2: %.5f\n", dBest.x1, dBest.x2)
		fmt.Printf("f(xBest) = %.5f\n\n", dBest.f())
	}
	fmt.Println("Final answer: ")
	fmt.Printf("xBest: %.5f\n", dBest)
	fmt.Printf("f(xBest) = %.5f\n", dBest.f())
}
