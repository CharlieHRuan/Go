package main

import (
	"fmt"
	"math/big"
)

func main() {

	//big.Int
	distance := new(big.Int)
	distance.SetString("30000000000000000000000000000000", 10)
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	fmt.Println(lightSpeed, secondsPerDay)
	fmt.Println(distance)
	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)
	fmt.Println(seconds)

	//homework
	// 86400->bit.Int
	var n1 = big.NewInt(86400)
	var n2 = new(big.Int)
	n2.SetString("86400", 10)
	fmt.Println(n1, n2)

	//常量是可以无类型
	// const Num = 230000000000000000000000000000000
	// fmt.Println(Num) //报错，因为编译的时候，默认会将无类型转为int

	const distanceStar = 236000000000000000
	fmt.Printf("大矮星距离%v光年\n", 236000000000000000/9.461/10e15)
}
