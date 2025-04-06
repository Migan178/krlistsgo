package main

import (
	"fmt"

	"github.com/Migan178/krlistsgo"
)

func main() {
	krbots := krlistsgo.New("", "") // 필요 없는 갑
	user, err := krbots.User("415135882006495242")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
