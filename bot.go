package krlistsgo

import (
	"encoding/json"
	"net/http"
	"time"
)

// BotFlags는 봇의 플래그 타입입니다.
type BotFlags int

// BotLib는 봇이 사용하는 라이브러리 타입입니다.
type BotLib string

// BotState는 한디리에서의 봇 상태 타입입니다.
type BotState string

// BotStatus는 디스코드에서의 봇 상태 타입입니다.
type BotStatus string

// BotCategory는 봇의 카테고리 타입입니다.
type BotCategory string

// Bot은 한디리 API에서 반환된 봇 데이터 타입입니다.
type Bot[T any] struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Tag         string        `json:"tag"`
	Avatar      string        `json:"avatar"`
	Owners      []T           `json:"owners"`
	Flags       BotFlags      `json:"flags"`
	Lib         BotLib        `json:"lib"`
	Prefix      string        `json:"prefix"`
	Votes       int           `json:"votes"`
	Servers     int           `json:"servers"`
	Shards      int           `json:"shards"`
	Intro       string        `json:"intro"`
	Desc        string        `json:"desc"`
	Web         string        `json:"web"`
	Git         string        `json:"git"`
	Url         string        `json:"url"`
	Discord     string        `json:"discord"`
	Category    []BotCategory `json:"Category"`
	Vanity      string        `json:"vanity"`
	Bg          string        `json:"bg"`
	Banner      string        `json:"banner"`
	Status      BotStatus     `json:"status"`
	State       BotState      `json:"state"`
	client      *http.Client  `json:"-"`
	identify    *Identify     `json:"-"`
	lastUpdated time.Time     `json:"-"`
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
func (k *KrLists) Bot(id string) (bot *Bot[User[string, string]], err error) {
	if k.CacheInterval != 0 {
		if data, ok := k.CachedData.Bots[id]; ok {
			if data.lastUpdated.Unix()-int64(k.CacheInterval) < int64(k.CacheInterval) {
				return data, nil
			}
			delete(k.CachedData.Bots, id)
		}
	}

	resp, err := get(k.Client, EndpointBots(id), &map[string]string{
		"Authorization": k.BotIdentify.Token,
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &bot)
	if err != nil {
		return
	}

	bot.Discord = "https://discord.gg/" + bot.Discord
	bot.client = k.Client
	bot.identify = k.BotIdentify
	bot.lastUpdated = time.Now()

	for _, owner := range bot.Owners {
		owner.Github = "https://github.com/" + owner.Github
	}

	k.CachedData.Bots[id] = bot
	return
}

// UpdateServers는 해당 봇의 서버 수를 업데이트합니다.
func (k *KrLists) UpdateServers(servers, shards int) error {
	if k.BotIdentify == nil {
		return ErrBotIdentifyIsNil
	}

	body := map[string]int{
		"servers": servers,
		"shards":  shards,
	}

	_, err := post(k.Client, EndpointBotUpdateServers(k.BotIdentify.ID), body, &map[string]string{
		"Authorization": k.BotIdentify.Token,
	})
	return err
}
