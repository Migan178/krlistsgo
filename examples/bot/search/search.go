package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New("", "") // 필요 없는 값
	data, err := k.SearchBots("관리", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
