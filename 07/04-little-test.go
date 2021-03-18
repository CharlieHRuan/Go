//小测试
//函数类型重写函数签名
// func drawTable(rows int,getRow func(row int)(string,string))

package main

import (
	"fmt"
	"strconv"
)

type simple func(row int) (string, string)

func drawTable(rows int, s simple) {
	str1, str2 := s(rows)
	fmt.Println(str1, str2, "行")
}

func getRow(row int) (string, string) {
	return "生成", strconv.Itoa(row)
}

func main() {
	drawTable(2, getRow)
}
 