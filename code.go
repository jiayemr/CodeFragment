package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rands = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {

	var lens = 100
	var arr [100]int
	for a := 0; a < lens; a++ {
		arr[a] = rands.Intn(9999)
	}
	str := arr[0:lens]

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str)-i-1; j++ {
			if str[j] > str[j+1] {
				tmp := str[j]
				str[j] = str[j+1]
				str[j+1] = tmp
			}
		}
	}

	for x := 0; x < len(str); x++ {
		fmt.Println(str[x])
	}

}
