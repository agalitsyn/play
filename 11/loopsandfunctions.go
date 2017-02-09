// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) float64 {
	result, delta := 1.0, 1.0

	for delta > 1e-15 {
		prev := result
		result = result - (math.Pow(result, result)-num)/(2*result)
		delta = math.Abs(result - prev)
	}

	return result
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}