package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	url, err := k.ServerWidget("704999866094452816", krlistsgo.WidgetServerVotes, krlistsgo.WidgetFlat, 1.0, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}
