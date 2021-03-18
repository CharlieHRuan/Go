package main

import "fmt"

func main() {
	const remainDays = 28
	const distance = 56000000

	var speed = distance / remainDays
	fmt.Printf("飞船行进速度是 %v km/h.", speed/24)
}
