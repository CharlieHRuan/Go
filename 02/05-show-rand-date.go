// 展示随机日期

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {

	for i := 0; i < 10; i++ {
		year, month, day := RandDate()
		fmt.Println(era, year, month, day)
	}
}

func RandDate() (int, int, int) {
	time.Sleep(time.Second)
	rand.Seed(time.Now().UnixNano())
	year := rand.Intn(2099-1700) + 1700
	var month, daysIsMonth = 0, 31
	month = rand.Intn(12) + 1
	daysIsMonth = 31
	switch month {
	case 2:
		daysIsMonth = CalcDays(year)
	case 4, 6, 9, 11:
		daysIsMonth = 30
	}
	day := rand.Intn(daysIsMonth) + 1
	return year, month, day
}

func CalcDays(year int) int {
	leapYear := year%4 == 0
	if leapYear {
		return 29
	}
	return 28
}
