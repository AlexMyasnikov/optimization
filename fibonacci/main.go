package main

import (
	"fmt"
	"math"
)

var fib = []float64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233}

func getS(N float64) int {
	for i, f := range fib {
		if float64(f) > N {
			return i
		}
	}
	return -1
}

func f(x float64) float64 {
	return 10 / (x * (4 - math.Pow(x, 2)))
}

func main() {
	a, b := 1.0, 1.5
	e := 0.01
	N := (b - a) / e
	S := getS(N)
	k := 1
	fmt.Printf("N: %.5f\n", N)
	fmt.Printf("S: %d\n", S)
	l := (b - a) / fib[S]
	x1 := a + l*fib[S-2]
	x2 := a - l*fib[S-2]
	f1 := f(x1)
	f2 := f(x2)
	for {
		k++
		if f1 < f2 {
			b = x2
			if k == S-1 {
				x2 = x1 + e
				f2 = f(x2)
				break
			} else {
				x2 = x1
				f2 = f1
				x1 = a + l*fib[S-1-k]
				f1 = f(x1)
			}
		} else {
			a = x1
			if k == S-1 {
				x2 = x1 + e
				f2 = f(x2)
				break
			} else {
				x1 = x2
				f1 = f2
				x2 = b - l*fib[S-1-k]
				f2 = f(x2)
			}
		}
		fmt.Printf("k: %d, x1: %.5f, x2: %.5f\n", k, x1, x2)
	}
	if f1 < f2 {
		b = x1
	} else {
		a = x1
	}
	x := (a + b) / 2
	R := f(x)
	fmt.Printf("x: %.5f, R: %.5f\n", x, R)
}
