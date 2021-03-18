package main

import "fmt"

func main() {
	c := 'y'
	c += 3
	c = CaesarCipher(c)
	fmt.Printf("%c", c)

	c = 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c", c)
}

func CaesarCipher(c rune) rune {
	//从头开始
	if c > 'z' {
		c -= 26
	}
	return c
}
