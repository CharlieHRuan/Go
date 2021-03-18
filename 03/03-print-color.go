package main

import "fmt"

func main() {
	var red, green, blue uint8 = 0, 22, 123
	fmt.Printf("%x %x %x\n", red, green, blue)

	//to color code
	fmt.Printf("#%02x%02x%02x\n", red, green, blue)
}
