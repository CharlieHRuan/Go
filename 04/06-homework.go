package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// cipherStr := "L fdph, L vdz, L frqtxhuhg"

	// for i := 0; i < len(cipherStr); i++ {
	// 	c := cipherStr[i]
	// 	if c >= 'a' && c <= 'z' {
	// 		c += 3
	// 		if c > 'z' {
	// 			c -= 26
	// 		}
	// 	}
	// 	if c >= 'A' && c <= 'Z' {
	// 		c += 3
	// 		if c > 'Z' {
	// 			c -= 26
	// 		}
	// 	}
	// 	fmt.Printf("%c\n", c)
	// }

	message := "Hola Estación Espacial Internacional"
	utf8.RuneCountInString(message)
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	// question := "¿Cómo estás?"

	// //len 计算字符串所占字节数
	// // fmt.Println(len(question), "bytes")
	// // fmt.Println(utf8.RuneCountInString(question), "runes")

	// // c, size := utf8.DecodeRuneInString(question)
	// // fmt.Printf("First rune: %c %v bytes.\n", c, size)

	// // range
	// // _丢弃标识符，如果不需要这个参数，就可以用这个替代
	// // for i, c := range question {
	// // var index = 0
	// for _, c := range question {
	// 	if c <= 'a' && c >= 'z' {
	// 		c += 13
	// 		if c > 'z' {
	// 			c -= 26
	// 		}
	// 	}
	// 	fmt.Printf("%c", c)
	// }
}
