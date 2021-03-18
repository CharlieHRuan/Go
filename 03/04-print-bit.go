package main

import (
	"fmt"
	"math"
)

func main() {
	var green uint8 = 255
	fmt.Printf("%08b\n", green)

	green++
	fmt.Printf("%08b\n", green)

	var max = math.MaxInt64
	fmt.Println(max)
}
