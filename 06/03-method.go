package main

import "fmt"

type kelvin float64
type celsius float64

func main() {
	//  声明新类型,增加代码可读性
	// type celsius float64

	// 温度，摄氏度
	var temperature celsius = 20

	// 类型不能混用
	var warmup float64
	// temperature += warmup	//报错，不能混用，是两种类型

	fmt.Println(temperature)

	k := 40
	fmt.Println("k", k.celsius())
}

//函数（函数名，参数类型，返回类型）
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

//方法（kelvin关联的方法,kelvin类上面有个方法Celsius）
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//错误：不能关联预声明类型
// func (k float64) celsius() celsius {
///....
// }
