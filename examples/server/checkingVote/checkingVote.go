package main

import (
	"flag"
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	token := flag.String("token", "", "server's koreanbots token")
	id := flag.String("id", "", "server's id")

	flag.Parse()

	k := krlistsgo.New().SetServerIdentify(*token, *id)
	data, err := k.CheckServerVote("415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
