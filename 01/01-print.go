package main

import "fmt"

func main() {
	//常量
	const lightSpeed = 111222
	// 变量
	var distance = 100000099
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}
