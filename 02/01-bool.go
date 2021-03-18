package main

import (
	"fmt"
	"strings"
)

func main() {
	// 包含
	var command = "cmd format"

	var result = strings.Contains(command, "fmt")

	fmt.Println("is contains :", result)

	//比较
	// > < = <= >= != ==

	//if

	var age = 18

	if age >= 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	// && 与 || 或

	// ! 取反

	switch command {
	case "go east":
		fmt.Println("命令格式化")
	case "cmd format":
		fmt.Println("未知命令")
		// fallthrough		//会继续执行下一个语句，类似c#
	case "cmd format1":
		fmt.Println("未知命令1")
	default:
		fmt.Println("默认输出")
	}

}
