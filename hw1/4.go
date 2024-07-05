package main

import (
	"fmt"
	"math"
)

func maxOfThree(x1, x2, x3 int) int {
	return (int)(math.Max(math.Max((float64)(x1), (float64)(x2)), (float64)(x3)))
}

func main() {
	fmt.Println(maxOfThree(2, 10, 3))
}
