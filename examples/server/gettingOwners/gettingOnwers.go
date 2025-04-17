package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New("", "") // 필요 없는 값
	owners, err := k.ServerOwners("909768169248395274")
	if err != nil {
		panic(err)
	}
	fmt.Println(owners)
}
