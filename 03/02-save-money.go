package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var target = 20.0
	var currentMoney = 0.0
	for currentMoney <= target {
		var temp = moneyType()
		currentMoney += temp
		fmt.Printf("本次存入金额：%.2f，总金额目前：%5.2f\n", temp, currentMoney)
	}
	fmt.Println("本次存入金额完成")
}

func moneyType() float64 {
	num := rand.Intn(3)
	result := 0.0
	switch num {
	case 0:
		result = 0.05
	case 1:
		result = 0.1
	case 2:
		result = 0.25
	}
	return float64(result)
}
