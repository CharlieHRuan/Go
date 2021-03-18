package main

import "fmt"

func main() {
	// 字符串字面值
	fmt.Println("nihao\n我是moumou")
	// 字符串原始字面值
	fmt.Println(`nihao \n 我是moumou`)
	fmt.Println(`nihao \n 
	我是moumou`)

	//类型别名
	// rune 是uint32的别名
	// byte 是uint8的别名
	// 自定义别名(b8表示uint8)
	type b8 = uint8

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang rune = 33
	var smile rune = 128515
	fmt.Println(smile)
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
	// 字符字面值也可以使用byte类型，但不是所有的
	var star byte = '*'
	fmt.Println(star)

	//可以给某个变量赋值不同的string，但是string本身是不可变的
	
}
