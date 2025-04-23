package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	data, err := k.Server("909768169248395274")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
