package main

import (
	"flag"

	"github.com/Migan178/krlistsgo"
)

func main() {
	token := flag.String("token", "", "discordbot's koreanbots token")
	id := flag.String("id", "", "discordbot's id")

	flag.Parse()

	k := krlistsgo.New().SetBotIdentify(*token, *id)

	err := k.UpdateServers(20, 1)
	if err != nil {
		panic(err)
	}
}
