package main

import "fmt"

func main() {
	var k kelvin = 294.0
	var c celsius

	c = kelvinToCelsius(k)
	c = k.celsius()

	fmt.Println(c)
}

type kelvin float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
