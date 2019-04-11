package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y, sum := 0, 1, 0
	fmt.Println(x)
	fmt.Println(y)
	return func() int {
		sum = x + y
		x = y
		y = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
