//一等函数 ：
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 函数赋给变量
type kelvin float64
type celsius float64

func fakeSensor() kelvin {
	rand.Seed(time.Now().UnixNano())
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func groundSensor() celsius {
	return 0
}

func main() {
	//声明一个函数变量
	// var sensor func() kelvin
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	//不能这么用，存在类型转换
	// sensor = groundSensor
}

// **************************************************
