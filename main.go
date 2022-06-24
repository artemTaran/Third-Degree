package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SumCubes(3))
}

func SumCubes(n int) (out int) {
	for i := 0; i < n; i++ {
		out += int(math.Pow(float64(i+1), 3))
	}
	return out
}
