package main

import "fmt"

func main() {
	var k kelvin = 294.0
	sensor := func() kelvin {
		return k
	}

	fmt.Println(sensor()) //294

	k++
	fmt.Println(sensor()) //295
}

type kelvin float64
