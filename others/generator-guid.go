package main

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"time"
)

func main() {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("错误")
	}
	uiid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uiid)

	fmt.Println(time.Unix(1616462554, 0).UTC().Add(time.Duration(f2i(1616463533*3600)) * time.Second))
}

func f2i(f float32) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}
