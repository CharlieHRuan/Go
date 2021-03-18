package main

//维吉尼亚算法
import (
	"fmt"
	"strconv"
)

func main() {
	key := "GOLANG"
	originStr := "CSOITEUIWUIZNSROCNKFD"
	decryptStr := Decrypt(key, originStr)
	fmt.Println(decryptStr)
	// decryptStr := "IGZIGKAWHUVFTGCOPTQTO"
	encryptStr := Encrypt(key, decryptStr)
	fmt.Println(encryptStr)
}

// 加密，密钥，待加密字符串
func Encrypt(key string, enStr string) string {
	result := ""
	idx := 0
	for _, o := range enStr {
		// var origStr = o
		if idx >= len(key) {
			idx = 0
		}
		val, _ := strconv.Atoi(fmt.Sprintf("%v", key[idx]))
		idx++
		var offset = val - 64 - 1
		if o >= 'A' && o <= 'Z' {
			o += rune(offset)
			if o > 'Z' {
				o -= 26
			}
		}
		result += string(o)
		// fmt.Printf("%c + %v,%c = %c\n", origStr, offset, val, o)
	}
	return result
}

// 解密，密钥，待解密字符串
func Decrypt(key string, deStr string) string {
	result := ""
	idx := 0
	for _, o := range deStr {
		// var origStr = o
		if idx >= len(key) {
			idx = 0
		}
		val, _ := strconv.Atoi(fmt.Sprintf("%v", key[idx]))
		idx++
		var offset = val - 64 - 1
		if o >= 'A' && o <= 'Z' {
			o -= rune(offset)
			if o < 'A' {
				o += 26
			}
		}
		result += string(o)
		// fmt.Printf("%c + %v,%c = %c\n", origStr, offset, val, o)
	}
	return result
}
