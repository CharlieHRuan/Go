//测试题

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//距火星距离
const distance = 62100000

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%-20v %-6v %-10v %5v\n", "Spaceline", "Days", "Trip Type", "Price")
	fmt.Printf("============================================\n")
	for i := 0; i < 100; i++ {
		GeneratorTrip()
	}
}

func GeneratorTrip() {
	var days, price, typeStr = GetRandDay()
	var company = GetRandSpaceLine()
	fmt.Printf("%-20v %-6v %-10v $%4v\n", company, days, typeStr, price)
}

func GetRandSpaceLine() string {
	result := ""
	line := rand.Intn(3) + 1
	switch line {
	case 1:
		result = "Space Adventures"
	case 2:
		result = "SpaceX"
	case 3:
		result = "Virgin Galactic"
	}
	return result
}

func GetRandDay() (int, int, string) {
	speed := rand.Intn(30-16) + 16
	price := rand.Intn(51-36) + 36
	s := distance / speed
	days := s / 60 / 60 / 24
	//类型，往返
	var types, str = GetTripType()
	days *= types
	price *= types
	return days, price, str
}

func GetTripType() (int, string) {
	var types, str = 0, ""
	types = rand.Intn(3-1) + 1
	if types == 1 {
		str = "One-way"
	} else {
		str = "Round-trip"
	}
	return types, str
}
