package main

import "fmt"

func main() {
	const lightSpeed = 111222
	var distance = 100000099
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}
