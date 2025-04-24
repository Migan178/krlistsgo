package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	data, err := k.SearchBots("한디리", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
