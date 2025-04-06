# KrlistsGo

- [한국 디스코드 리스트](https://koreanbots.dev)의 **비공식** Golang API 래퍼.
- [본래 다른 프로젝트(제작자 다름)](https://pkg.go.dev/github.com/simsimler/gokoreanbots)이 있있지만, 해당 패키지는 더이상 개발이 되지 않는 관계로 새로이 작성하였습니다.

## 사용법

- 더 많은 예제는 [여기서 확인 가능.](https://github.com/Migan178/krlistsgo/tree/main/examples)

### 설치

```sh
go get github.com/Migan178/krlistsgo
```

### 서버수 업데이트

```go
// 나중에 discordgo를 통한 자동 업데이트 지원 예정
package main

import "github.com/Migan178/krlistsgo"

func main() {
	// 서버 수
	servers := 20
	// 샤드 수
	shards := 1
	krlists := krlistsgo.New("한디리 토큰", "봇의 디스코드 아이디")
	err := krlists.UpdateServers(servers,shards)
	// 예외 처리
	if err !=nil {
		panic(err)
	}
}
```
