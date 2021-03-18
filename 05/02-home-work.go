package main

import "fmt"

func main() {

	flag := true

	printFlag := -1
	if flag {
		printFlag = 1
	} else {
		printFlag = 0
	}

	fmt.Println(printFlag)

	boolStr := "false"

	fmt.Println(fmt.Sprintf("%v", "yes"))

	if boolStr == fmt.Sprintf("%v", true) || boolStr == "yes" || boolStr == "1" {
		fmt.Println(true)
	} else if boolStr == fmt.Sprintf("%v", false) || boolStr == "no" || boolStr == "0" {
		fmt.Println(false)
	} else {
		fmt.Println("error")
	}

}
