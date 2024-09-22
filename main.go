package main

import (
	"fmt"
	"time"
)

func main() {
	s := CreateJwtToken("cdbb", time.Second*3600)
	fmt.Printf("Hello and welcome, %s!\n", s)

	claim, err := ParseToken(s)
	if err == nil {
		fmt.Println(claim)
	}

}
