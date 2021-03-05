package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println("Hello")

	fmt.Print("Hello,23")
	fmt.Println("Hello,34")

	fmt.Printf("体重是 %v 磅", 149.0*0.3783)
	fmt.Printf("体重是 %v 磅\n", 149.0*0.3783)

	fmt.Printf("体重在%v是 %v 磅\n", "火星", 149.0*0.3783)

	//对齐文本
	//-向右填充，+向左填充
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
	/*
		SpaceX          $  94
		Virgin Galactic $ 100
	*/
}
