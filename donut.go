package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	A = 1.0
	B = 1.0
)

var x, y, o, N int
var cA, sA, cB, sB, ct, st, sp, cp, D, t, h float64
var chars = strings.Split(".,-~:;=!*#$@", "")

func donut(z [1760]float64, b [1760]string) {
	A += 0.07
	B += 0.03

	cA = math.Cos(A)
	sA = math.Sin(A)
	cB = math.Cos(B)
	sB = math.Sin(B)

	for j := 0.0; j < 6.28; j += 0.07 {
		ct = math.Cos(j)
		st = math.Sin(j)
		for i := 0.0; i < 6.28; i += 0.02 {
			sp = math.Sin(i)
			cp = math.Cos(i)
			h = ct + 2

			D = 1 / (sp*h*sA + st*cA + 5)
			t = sp*h*cA - st*sA

			x = 0 | int(40+30*D*(cp*h*cB-t*sB))
			y = 0 | int(12+15*D*(cp*h*sB+t*cB))
			o = x + 80*y
			N = 0 | int(8*((st*sA-sp*ct*cA)*cB-sp*ct*sA-st*cA-cp*ct*sB))

			if 22 > y && y > 0 && x > 0 && 80 > x && D > z[o] {
				z[o] = D

				if N > 0 {
					b[o] = chars[N]
				} else {
					b[o] = chars[0]
				}
			}

			fmt.Fprintf(os.Stdout, "\x1b[H")
		}

		fmt.Fprintf(os.Stdout, strings.Join(b[:], ""))
	}
}

func main() {
	var z [1760]float64
	var b [1760]string

	for k := 0; k < 1760; k++ {
		z[k] = 0
		if k%80 == 79 {
			b[k] = "\n"
		} else {
			b[k] = " "
		}
	}

	for {
		fmt.Print("\x1b[2J")
		donut(z, b)
	}
}
