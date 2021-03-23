package main

import "fmt"

func main() {
	var arr [4]string

	arr[0] = "Mars"
	arr[1] = "Earth"

	fmt.Println(arr[2] == "")
	// i := 8
	// fmt.Println(arr[i])

	//复合字面值初始化数组
	// dwarf := [5]string{"C", "P", "A", "A", "D"}
	//... 长度固定，不可变
	dwarfs := [...]string{"C", "P", "A", "A", "D"}
	fmt.Println(dwarfs)

	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
	fmt.Println("=====================")
	for i, v := range dwarfs {
		fmt.Println(i, v)
	}

	// 数组赋值给新变量，还是传递给函数，都会产生一个完整的拷贝

	// 数组也是一种值，函数通过值传递接受参数，数组作为函数参数（也是整个复制一遍），效率非常低

	// 二维数组
 
}
