package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	// var count = 0
	// for count = 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }
	// fmt.Println(count)

	// 短声明
	// for count := 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }

	// if count := rand.Intn(3); count == 0 {
	// 	fmt.Println("零")
	// } else if count == 2 {
	// 	fmt.Println("二")
	// } else {
	// 	fmt.Println("一")
	// }

	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("零")
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	case 4:
		fmt.Println("四")
	default:
		fmt.Println(num)
	}

	//短声明不能用来声明package作用域的变量
	// 下面的era 就不能使用短声明  
}

var era = "AD"
