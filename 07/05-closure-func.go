//闭包，匿名函数
//匿名函数，go里面称为函数字面值
package main

import "fmt"

func main() {
	f()

	g := func(message string) {
		fmt.Println(message)
	}
	g("另外一种匿名函数")

	//只用一次，声明后立即执行，小括号相当于执行函数
	func() {
		fmt.Println("第三种匿名函数")
	}()
}

// 这就是一个匿名函数，函数没有名称
var f = func() {
	fmt.Println("匿名函数")
}
