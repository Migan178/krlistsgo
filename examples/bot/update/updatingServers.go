package main

import (
	"flag"

	"github.com/Migan178/koreanbotsgo"
)

func main() {
	token := flag.String("token", "", "discordbot's koreanbots token")
	id := flag.String("id", "", "discordbot's id")

	flag.Parse()

	krbots := koreanbotsgo.New(*token, *id)

	err := krbots.UpdateServers(20, 1)
	if err != nil {
		panic(err)
	}
}
