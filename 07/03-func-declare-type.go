// 函数声明类型
// 函数声明类型可以精简和明确调用者的代码

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func main() {
	measureTemperature(3, fakeSensor)
}

func measureTemperature(samples int, s sensor) {
	// func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		// k := sensor()
		k := s()
		fmt.Printf("%v °K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
