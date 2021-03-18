package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	c = c*9.0/5.0 + 32.0
	return c
}

func kelvinToFarenheit(k float64) float64 {
	k -= 459.67
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K is ", celsius, "°C\n")

	kelvin = 223.0
	celsius = kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K is ", celsius, "°C\n")

	celsius = 36.5
	f := celsiusToFahrenheit(celsius)
	fmt.Print(celsius, "°C is ", f, "°F\n")

	kelvin = 0
	f = kelvinToFarenheit(kelvin)
	fmt.Print(celsius, "°C is ", f, "°F\n")
}
