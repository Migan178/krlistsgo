package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	url, err := k.BotWidget("704999866094452816", krlistsgo.WidgetBotVotes, krlistsgo.WidgetFlat, 1.0, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
