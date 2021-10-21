package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 5.2
	var b float64 = 1.1
	var c = (math.Sin(a) + math.Cos(b)) / math.Cos(5.6)
	fmt.Printf("Result: %v\n", c)
}
