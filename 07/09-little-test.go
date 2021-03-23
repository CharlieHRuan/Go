//温度表

package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64
type print func(float64)

func main() {
	Splite()
	// PrintFormat("°C", "°F")
	PrintFormat("°F", "°C")
	Splite()
	// var c celsius = -40
	var f fahrenheit = -40
	// for c <= 100 {
	// 	Calc(float64(c), "c2f")
	// 	c += 5
	// }

	for f <= 100 {
		Calc(float64(f), "f2c")
		f += 5
	}
	Splite()
}

func Calc(temp float64, flag string) {
	switch flag {
	case "c2f":
		drawTable(celsiusToFahrenheit, temp)
	case "f2c":
		drawTable(FahrenheitTocelsius, temp)
	}

}

func drawTable(p print, c float64) {
	p(c)
}

//分割线
func Splite() {
	fmt.Printf("=========================\n")
}

// 摄氏度转华氏度单位
func celsiusToFahrenheit(c float64) {
	f := c*1.8 + 32
	tempF := fmt.Sprintf("%v", f)
	tempC := fmt.Sprintf("%v", c)
	PrintFormat(tempC, tempF)
	// return fmt.Sprintf("|%-11v|%-11v|", c, f)
}

//华氏度转摄氏度单位
func FahrenheitTocelsius(f float64) {
	c := (f - 32) / 1.8
	tempF := fmt.Sprintf("%v", f)
	tempC := fmt.Sprintf("%.2f", c)
	PrintFormat(tempF, tempC)
	// return fmt.Sprintf("|%-11v|%-11v|", f, c)
}

//格式化打印
func PrintFormat(str1 string, str2 string) {
	fmt.Printf("|%-11v|%-11v|\n", str1, str2)
}
