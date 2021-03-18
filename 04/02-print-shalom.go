package main

import "fmt"

func main() {
	str := "shalom"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}
