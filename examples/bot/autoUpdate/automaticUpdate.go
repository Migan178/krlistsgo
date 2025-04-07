package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Migan178/krlistsgo"
	"github.com/bwmarrin/discordgo"
)

func main() {
	kToken := flag.String("kToken", "", "discordbot's koreanbots token")
	dToken := flag.String("dToken", "", "discordbot's discord token")

	flag.Parse()

	dg, err := discordgo.New("Bot " + *dToken)
	if err != nil {
		panic(err)
	}

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	// 해당 함수를 사용함으로써 반환되는 *KrLists 값은 필요 없어서 값 할당 안함.
	// 봇이 시작된 (= 로그인된)뒤에 해당 함수를 사용해줘야함.
	krlistsgo.NewKrListsWithDiscordGo(dg, *kToken, true)

	defer func() {
		dg.Close()
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
