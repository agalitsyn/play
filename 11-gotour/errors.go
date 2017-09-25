// https://tour.golang.org/methods/20
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (self ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(self))
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegativeSqrt(num)
	}

	result := 1.0
	for i := 1; i < 1000; i++ {
		result = result - (((result * result) - num) / 2 * result)
	}

	return result, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}