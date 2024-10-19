package main

import (
	"fmt"
	"time"
)

func main() {
	key := "0123456789ABCDEF"
	s := CreateToken("cdbb", time.Second*3600, key)
	fmt.Printf("Hello and welcome, %s!\n", s)

	claim, err := ParseToken(s, key)
	if err == nil {
		fmt.Println(claim)
	}
}
