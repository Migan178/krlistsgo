package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	krbots := krlistsgo.New("", "") //필요 없는 값
	server, err := krbots.Server("909768169248395274")
	if err != nil {
		panic(err)
	}
	fmt.Println(server)
}
