package main

import "fmt"

//变量常量
func main() {
	//变量、常量
	const lightSpeed = 111222
	var distance = 100000099
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	//多个变量赋值
	//未使用变量会报错
	var a, b = 32, 60
	var (
		c = 3
		d = 4
	)
	const e, f = 43, 44

	// 赋值运算符
	var w = 1223
	w = 3
	w *= 3

	// 自增运算
	var age = 12 + 1
	age += 1
	age++
	//报错 ++age

}
