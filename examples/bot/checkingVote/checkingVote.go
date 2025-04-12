package main

import (
	"flag"
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	token := flag.String("token", "", "discordbot's koreanbots token")
	id := flag.String("id", "", "discordbot's id")

	flag.Parse()

	krLists := krlistsgo.New(*token, *id)
	data, err := krLists.CheckBotVote("415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
