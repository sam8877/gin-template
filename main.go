package main

import (
	"fmt"
	"time"
)

func main() {
	key := "0123456789ABCDEF"
	s, err := CreateToken("cdbb", time.Second*1, key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)

	claim, err := ParseToken(s, key)
	if err == nil {
		fmt.Println(claim)
	}

	time.Sleep(time.Second * 2)

	exped := claim.Expired()
	if exped {
		s, err := CreateToken("cdbb", time.Second*1, key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(s)
	}

}
