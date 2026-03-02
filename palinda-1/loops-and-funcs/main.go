package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	n := 0.0
	i := 1
	off := 0.00000001
	/*for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z*/
	for {
		z -= (z*z - x) / (2 * z)

		if z-n > 0 && z-n < off {
			println("number of iteratiosn", i)
			return z

		} else if z-n < 0 && n-z < off {
			println(i)
			return z
		}
		n = z
		i++
	}

}
func main() {
	x := 2.0
	fmt.Println("start")
	fmt.Println("Sqrt function: ", Sqrt(x))
	fmt.Println("math.sqrt:", math.Sqrt(x))
	fmt.Println("diff: ", math.Sqrt(x)-Sqrt(x))
}
