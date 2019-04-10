package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z, y := 1.0, x
	for i := 1; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
		if (y - z) < 1e-15 {
			fmt.Printf("times: %d\n", i)
			return z
		} else {
			y = z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
