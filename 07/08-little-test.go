//修改程序

//声明一个变量，将其作为calibrate函数的offset实参，而不是使用字面值数字5，
// 此后，即使修改变量，调用sensor结果也仍然是5，
// 这是因为，offset形参接受的实参副本而不是引用，也就是按值传递

package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type fKelvin float64

type sensor func() kelvin
type fsensor func() fKelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func fcalibrate(f fsensor, offset fKelvin) fsensor {
	return func() fKelvin {
		return f() + offset
	}

}
func fakeSensor() fKelvin {
	return fKelvin(rand.Intn(151) + 150)
}

func main() {
	var k kelvin = 5
	sensor := calibrate(realSensor, k)
	fmt.Println(sensor())
	k++
	fmt.Println(sensor())

	fSensor := fcalibrate(fakeSensor, 5)
	fmt.Println(fSensor())
}
