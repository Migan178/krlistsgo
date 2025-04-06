package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	krbots := krlistsgo.New("", "") // 필요 없는 값
	bot, err := krbots.Bot("704999866094452816")
	if err != nil {
		panic(err)
	}
	fmt.Println(bot)
}
