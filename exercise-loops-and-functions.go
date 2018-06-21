package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z=x
	for math.Abs(z*z-x)>0.0001{
		z -= (z*z - x) / (2*z)	
	}
	return z
	
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}