package main

import (
	"fmt"
	"log"

	"github.com/Migan178/koreanbotsgo"
)

func main() {
	krbots := koreanbotsgo.New("", "") // 필요 없는 값
	bot, err := krbots.Bot("704999866094452816")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bot)
}
