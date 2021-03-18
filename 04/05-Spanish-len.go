package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿Cómo estás?"

	//len 计算字符串所占字节数
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes.\n", c, size)

	// range
	// _丢弃标识符，如果不需要这个参数，就可以用这个替代
	// for i, c := range question {
	for _, c := range question {
		fmt.Printf("%c\n", c)
	}
}
