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
	kToken := flag.String("kToken", "", "봇의 디스코드 토큰")
	dToken := flag.String("dToken", "", "봇의 한디리 토큰")

	flag.Parse()

	dg, err := discordgo.New("Bot " + *dToken)

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	krlistsgo.NewKrListsWithDiscordGo(dg, *kToken, true)

	defer dg.Close()

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
