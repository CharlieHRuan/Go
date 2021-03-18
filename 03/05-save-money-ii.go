package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var target = 2000
	var currentMoney = 0
	for currentMoney <= target {
		var temp = moneyType()
		currentMoney += (temp)
		fmt.Printf("本次存入金额：$%2v美分，总金额目前：$%5.2f美元\n", temp, float64(currentMoney)/100)
	}
	fmt.Println("本次存入金额完成")
}

func moneyType() int {
	num := rand.Intn(3)
	result := 0
	switch num {
	case 0:
		result = 5
	case 1:
		result = 10
	case 2:
		result = 25
	}
	return result
}
