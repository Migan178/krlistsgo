package krlistsgo

import (
	"encoding/json"
	"net/http"
	"time"
)

// ServerFlags는 서버의 플래그 타입입니다.
type ServerFlags int

// ServerState는 서버의 상태 타입입니다.
type ServerState string

// ServerCategory는 서버의 카테고리 타입입니다.
type ServerCategory string

// Server는 한디리 API에서 반환된 서버 데이터입니다.
type Server[T any] struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Icon        string           `json:"icon"`
	Owner       T                `json:"owner"`
	Flags       ServerFlags      `json:"flags"`
	Votes       int              `json:"votes"`
	Members     int              `json:"members"`
	BoostTier   int              `json:"boostTier"`
	Intro       string           `json:"intro"`
	Desc        string           `json:"desc"`
	Category    []ServerCategory `json:"category"`
	Invite      string           `json:"invite"`
	Emojis      []*ServerEmoji   `json:"emojis"`
	Vanity      string           `json:"vanity"`
	Bg          string           `json:"bg"`
	Banner      string           `json:"banner"`
	State       ServerState      `json:"state"`
	client      *http.Client     `json:"-"`
	identify    *Identify        `json:"-"`
	lastUpdated time.Time        `json:"-"`
}

// ServerEmoji는 서버의 이모지 구조체입니다.
type ServerEmoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// 서버의 플래그입니다.
const (
	ServerNone            ServerFlags = 0 << 0
	ServerOfficial        ServerFlags = 1 << 0
	ServerKrBotsVerified  ServerFlags = 1 << 2
	ServerKrBotsPartner   ServerFlags = 1 << 3
	ServerDiscordVerified ServerFlags = 1 << 4
	ServerDiscordPartner  ServerFlags = 1 << 5
)

// 서버의 상태입니다.
const (
	ServerOk          ServerState = "ok"
	ServerReported    ServerState = "reported"
	ServerBlocked     ServerState = "blocked"
	ServerUnReachable ServerState = "unreachable"
)

// 서버의 카테고리입니다.
const (
	ServerCommunity ServerCategory = "커뮤니티"
	ServerITScience ServerCategory = "IT & 과학"
	ServerBot       ServerCategory = "봇"
	ServerAmity     ServerCategory = "친목"
	ServerMusic     ServerCategory = "음악"
	ServerEducation ServerCategory = "교육"
	ServerDate      ServerCategory = "연애"
	ServerGame      ServerCategory = "게임"
	ServerOverwatch ServerCategory = "오버워치"
	ServerLOL       ServerCategory = "리그 오브 레전드"
	ServerPUBG      ServerCategory = "배틀그라운드"
	ServerMinecraft ServerCategory = "마인크래프트"
)

// Server의 정보를 갖고옵니다.
func (k *KrLists) Server(id string) (server *Server[User[string, string]], err error) {
	if k.CacheInterval != 0 {
		if data, ok := k.CachedData.Servers[id]; ok {
			if data.lastUpdated.Unix()-int64(k.CacheInterval) < int64(k.CacheInterval) {
				return data, nil
			}
			delete(k.CachedData.Servers, id)
		}
	}

	resp, err := get(k.Client, EndpointServers(id), &map[string]string{
		"Authorization": k.ServerIdentify.Token,
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &server)
	if err != nil {
		return
	}

	server.Invite = "https://discord.gg/" + server.Invite
	server.Owner.Github = "https://github.com/" + server.Owner.Github
	server.client = k.Client
	server.identify = k.ServerIdentify
	server.lastUpdated = time.Now()

	k.CachedData.Servers[id] = server
	return
}

// ServerOwners는 서버의 관리자들을 블러옵니다.
func (k *KrLists) ServerOwners(id string) (owners []User[string, string], err error) {
	resp, err := get(k.Client, EndpointServerOwners(id), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &owners)
	return
}
