package main

import (
	"flag"
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	token := flag.String("token", "", "봇의 한디리 토큰")
	id := flag.String("id", "", "봇의 디스코드 아이디")

	flag.Parse()

	k := krlistsgo.New().SetBotIdentify(*token, *id)
	data, err := k.CheckBotVote("415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
