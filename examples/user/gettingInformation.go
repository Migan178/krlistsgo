package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	k := krlistsgo.New()
	user, err := k.User("415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
