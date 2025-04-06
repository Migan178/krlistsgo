package koreanbotsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// BotFlags는 봇의 플래그입니다.
type BotFlags int

// Botlib는 봇이 사용하는 라이브러리입니다.
type BotLib string

// BotState는 한디리에서의 봇 상태입니다.
type BotState string

// BotStatus는 디스코드에서의 봇 상태입니다.
type BotStatus string

// BotCategory는 봇의 카테고리입니다.
type BotCategory string

// Bot은 한디리 API에서 반환된 봇 데이터입니다.
type Bot[T any] struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Tag      string        `json:"tag"`
	Avatar   string        `json:"avatar"`
	Owners   []User[T, T]  `json:"owners"`
	Flags    BotFlags      `json:"flags"`
	Lib      BotLib        `json:"lib"`
	Prefix   string        `json:"prefix"`
	Votes    int           `json:"votes"`
	Servers  int           `json:"servers"`
	Shards   int           `json:"shards"`
	Intro    string        `json:"intro"`
	Desc     string        `json:"desc"`
	Web      string        `json:"web"`
	Git      string        `json:"git"`
	Url      string        `json:"url"`
	Discord  string        `json:"discord"`
	Category []BotCategory `json:"Category"`
	Vanity   string        `json:"vanity"`
	Bg       string        `json:"bg"`
	Banner   string        `json:"banner"`
	Status   BotStatus     `json:"status"`
	State    BotState      `json:"state"`
}

// 봇의 플래그입니다.
const (
	BotNone            BotFlags = 0 << 0
	BotOfficial        BotFlags = 1 << 0
	BotKrBotsVerified  BotFlags = 1 << 2
	BotPartner         BotFlags = 1 << 3
	BotDiscordVerified BotFlags = 1 << 4
	BotPremium         BotFlags = 1 << 5
	BotHackathon       BotFlags = 1 << 6
)

// 봇이 사용하는 라이브러리입니다.
const (
	LibDiscordJs     BotLib = "discord.js"
	LibEris          BotLib = "Eris"
	LibDiscordPy     BotLib = "discord.py"
	LibDiscordCr     BotLib = "discordcr"
	LibNyxx          BotLib = "Nyxx"
	LibDiscordDotNet BotLib = "discord.net"
	LibDSharpPlus    BotLib = "DSharpPlus"
	LibNostrum       BotLib = "Nostrum"
	LibCoxir         BotLib = "coxir"
	LibDiscordGo     BotLib = "DiscordGo"
	LibDiscord4J     BotLib = "Discord4J"
	LibJavacord      BotLib = "Javacord"
	LibJDA           BotLib = "JDA"
	LibDiscordia     BotLib = "Discordia"
	LibRestCord      BotLib = "RestCord"
	LibYasmin        BotLib = "Yasmin"
	LibDisco         BotLib = "disco"
	LibDiscordrb     BotLib = "discordrb"
	LibSerenity      BotLib = "serenity"
	LibSwiftDiscord  BotLib = "SwiftDiscord"
	LibSword         BotLib = "Sword"
	LibOther         BotLib = "기타"
	LibPrivate       BotLib = "비공개"
)

// 한디리에서의 봇 상태입니다.
const (
	BotOk       BotState    = "ok"
	BotReported BotState    = "reported"
	BotBlocked  BotState    = "blocked"
	BotPrivate  BotState    = "private"
	BotArchived ServerState = "archived"
)

// 디스코드에서의 봇 상태입니다.
const (
	BotOnline    BotStatus = "online"
	BotIdle      BotStatus = "idle"
	BotDnD       BotStatus = "dnd"
	BotStreaming BotStatus = "streaming"
	BotOffline   BotStatus = "offline"
)

// Bot의 정보를 갖고옵니다.
func (k *Koreanbots) Bot(id string) (bot *Bot[string], err error) {
	var data ResponseBody
	req, err := http.NewRequest(http.MethodGet, API_URL+"/bots/"+id, nil)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := k.Client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	if data.Code != 200 {
		err = errors.New(fmt.Sprintf("Http Status Code: %d, Message: %s", data.Code, string(data.Message)))
		return
	}

	err = json.Unmarshal(data.Data, &bot)
	if err != nil {
		return
	}
	return
}
