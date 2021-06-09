package main

import (
	"fmt"
	"math"
)

func main() {
	// floating point follow standard IEEE754 http://mathcenter.oxford.emory.edu/site/cs170/ieee754/
	fmt.Println("max float32:", math.MaxFloat32)
	fmt.Println("max float64:", math.MaxFloat64)

	var f float64
	fmt.Println(f)
	fmt.Println(f/0) // NaN (not a number) maybe 0.0000000000000000001/0
	fmt.Println(1/f) // Inf (Infinity) maybe 1/0.00000000000000001

	f = 3.2
	fmt.Printf("f is %.2f\n",f)
}
