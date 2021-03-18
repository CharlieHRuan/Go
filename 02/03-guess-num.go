package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var resultNum = RandInt64(1, 100)
	fmt.Println("数字是：", resultNum)
	var guessnum = RandInt64(1, 100)
	fmt.Println("猜数字是：", guessnum)

	var low, high = 1, 100
	for guessnum != resultNum {
		if guessnum < resultNum {
			fmt.Println("数字太小了", guessnum)
			low = int(guessnum)
			guessnum = RandInt64(int64(low), int64(high))
		} else if guessnum > resultNum {
			fmt.Println("数字太大了", guessnum)
			high = int(guessnum)
			guessnum = RandInt64(int64(low), int64(high))
		} else if guessnum == resultNum {
			break
		}
	}
	fmt.Println("数字猜中了", guessnum)
}

func RandInt64(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
