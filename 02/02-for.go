package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10

	// 循环
	for count > 0 {
		fmt.Println(count)
		//time.Second 1秒
		//睡眠1秒继续执行
		time.Sleep(time.Second)
		count--
	}

	fmt.Println("Liftoff!")
}
