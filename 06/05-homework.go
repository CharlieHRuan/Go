//三种温度之间的转换方法

package main

import "fmt"

func main() {

	var c celsius = 20
	var k kelvin = 20
	var f fahrenheit = 20

	fmt.Printf("%v摄氏度->%v华氏度 ->%v开氏度\n", c, c.fahrenheit(), c.kelvin())
	fmt.Printf("%v华氏度->%v摄氏度 ->%v开氏度\n", f, f.celsius(), f.kelvin())
	fmt.Printf("%v开氏度->%v华氏度 ->%v摄氏度\n", k, k.fahrenheit(), k.celsius())

}

//声明三种类型
type celsius float64
type kelvin float64
type fahrenheit float64

// 开尔文转摄氏度
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// 开尔文转华氏度
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k-273.15)*1.8 + 32
}

//摄氏度转开尔文
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// 摄氏度转华氏度
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(1.8*c) + 32
}

// 华氏度转开尔文

func (f fahrenheit) kelvin() kelvin {
	return kelvin((f-32)/1.8 + 273.15)
}

// 华氏度转摄氏度
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) / 1.8)
}
