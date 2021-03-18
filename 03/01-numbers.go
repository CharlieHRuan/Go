package main

import (
	"fmt"
	"reflect"
)

func main() {

	// var answer float64 = 34 //使用整数声明浮点类型，去掉float就是整数类型
	answer := 42.0
	fmt.Println(reflect.TypeOf(answer)) //float64
	fmt.Printf("%T\n", answer)          //float64

	var price float64
	fmt.Println(price)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.2f\n", third)
	fmt.Printf("%10.2f\n", third)
	fmt.Printf("%010.2f\n", third)
	fmt.Printf("%4.2f\n", third)

	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")
	fmt.Print((celsius*9.0/5.0)+32, "° F\n")

}
