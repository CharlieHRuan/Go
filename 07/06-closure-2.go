//因为函数字面值需要保留外部作用域的变量引用，所以函数字面值都是闭包的

// 闭包：由于匿名函数封闭并包围作用域中的变量而得名
package main

import "fmt"

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		//保留（捕获）外部作用域的变量
		return s() + offset
	}
}

func main() {
	//虽然calibrate已经返回了，但是sensor 仍然能访问这两个变量
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
