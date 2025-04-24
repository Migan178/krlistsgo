# KrlistsGo

- [한국 디스코드 리스트](https://koreanbots.dev)의 **비공식** Golang API 래퍼.
- [본래 다른 프로젝트(제작자 다름)](https://pkg.go.dev/github.com/simsimler/gokoreanbots)이 있있지만, 해당 패키지는 더이상 개발이 되지 않는 관계로 새로이 작성하였습니다.

## 사용법

- 더 많은 예제는 [여기서 확인 가능합니다.](https://github.com/Migan178/krlistsgo/tree/main/examples)

### 설치

```sh
go get github.com/Migan178/krlistsgo
```

### 서버수 업데이트

#### 자동 업데이트

```go
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/Migan178/krlistsgo"
)

func main() {
	// 봇의 객체 생성
	dg, err := discordgo.New("Bot " + "디스코드 토큰")

	// 디스코드에 봇 로그인
	err = dg.Open()
	// 예외 처리
	if err != nil {
		panic(err)
	}

	// 해당 함수를 사용함으로써 반환되는 *KrLists 값은 필요 없어서 값 할당 안함.
	// 봇이 시작된 (= 로그인된)뒤에 해당 함수를 사용해줘야함. (여기서는 dg.Open() 이후에)
	// 해당 함수에서 3번째 인자는 서버수를 자동으로 업데이트 하도록 설정함.
	// 자동 업데이트 간격은 10분임.
	krlistsgo.NewKrListsWithDiscordGo(dg, "한디리 토큰", true)

	// 봇이 꺼지면 자동적으로 디스코드와의 연결을 끊음.
	defer dg.Close()

	// discordgo는 프로세스가 안 꺼질려면 채널을 만들어서 기달려야함.
	sc := make(chan os.Signal, 1)

	// 만약 프로세스가 꺼지면 sc 채널에 신호를 전송함.
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// 프로세스가 끝났다는 신호가 들어오기전까지 프로세스를 계속 유지하다가 신호를 받으면 꺼짐.
	// 이 신호 값을 받은 이후에는 defer dg.Close()가 실행됨.
	<-sc
}
```

#### 수동 업데이트

```go
package main

import "github.com/Migan178/krlistsgo"

func main() {
	// 서버 수
	servers := 20

	// 샤드 수
	shards := 1

	krlists := krlistsgo.New().SetBotIdentify("한디리 토큰", "봇의 디스코드 아이디")

	// 서버 수 업데이트
	err := krlists.UpdateServers(servers, shards)
	// 예외 처리
	if err !=nil {
		panic(err)
	}
}
```
