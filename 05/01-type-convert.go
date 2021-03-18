package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//字符串不能直接＋数字
	// 整数浮点数不能混用
	// fmt.Println("10" - 1)

	//环绕行为也会在类型转换中发生

	v := 42
	if v >= 0 && v <= math.MaxUint8 {
		v8 := uint8(v)
		fmt.Println("Convert:", v8)
	}

	//数字转字符串，必须有对应得code point 才行，比如65是A
	// 没有对应codepoint没关系，不会报错，但是不会转换成有意义的字符串
	fmt.Println(string(1))

	// 拼接数字+字符串
	fmt.Println("输出：" + strconv.Itoa(99999999999999999) + "个字符")
	res, err := strconv.Atoi("G")
	if err != nil {
		fmt.Println("输出异常,", err)
	}
	fmt.Println(res)
	str := fmt.Sprintf("输出： %v 个字符", 12)
	fmt.Println(str)
}
