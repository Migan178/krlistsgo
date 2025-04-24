package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	data, err := k.Bot("704999866094452816")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
