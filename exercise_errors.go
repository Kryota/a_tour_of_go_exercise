package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z, y := 1.0, x
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	} else {
		for i := 1; i < 10; i++ {
			z -= (z*z - x) / (2 * z)
			if (y - z) < 1e-15 {
				fmt.Printf("times: %d\n", i)
				return z, nil
			} else {
				y = z
			}
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
