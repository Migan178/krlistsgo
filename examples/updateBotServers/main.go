package main

import (
	"flag"

	"github.com/Migan178/krlistsgo"
)

func main() {
	token := flag.String("token", "", "봇의 한디리 토큰")
	id := flag.String("id", "", "봇의 디스코드 아이디")

	flag.Parse()

	servers := 20
	shards := 1

	k := krlistsgo.New().SetBotIdentify(*token, *id)
	err := k.UpdateServers(servers, shards)
	if err != nil {
		panic(err)
	}
}
