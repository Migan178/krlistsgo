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

	krLists := krlistsgo.New("", "") // 필요 없는 값
	data, err := krLists.CheckServerVote(*token, *id, "415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
