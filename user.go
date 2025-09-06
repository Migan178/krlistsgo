package krlistsgo

import (
	"encoding/json"
	"time"
)

// UserFlags는 유저의 플래그 타입입니다.
type UserFlags int

// User는 한디리 API에서 반환된 유저 데이터입니다.
type User[T, V any] struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Tag         string    `json:"tag"`
	Github      string    `json:"github"`
	Flags       UserFlags `json:"flags"`
	Bots        []T       `json:"bots"`
	Servers     []V       `json:"servers"`
	lastUpdated time.Time `json:"-"`
}

// 유저의 플래그입니다.
const (
	UserNone          UserFlags = 0 << 0
	UserAdministrator UserFlags = 1 << 0
	UserBugHunter     UserFlags = 1 << 1
	UserReviewer      UserFlags = 1 << 2
	UserPremium       UserFlags = 1 << 3
)

// User의 정보를 갖고옵니다.
func (k *KrLists) User(id string) (user *User[Bot[string], Server[string]], err error) {
	var token string

	if k.BotIdentify != nil {
		token = k.BotIdentify.Token
	} else {
		token = k.ServerIdentify.Token
	}

	if k.CacheInterval != 0 {
		if data, ok := k.CachedData.Users[id]; ok {
			if data.lastUpdated.Unix()-int64(k.CacheInterval) < int64(k.CacheInterval) {
				return data, nil
			}
			delete(k.CachedData.Users, id)
		}
	}

	resp, err := get(k.Client, EndpointUsers(id), &map[string]string{
		"Authorization": token,
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &user)
	if err != nil {
		return
	}

	user.Github = "https://github.com/" + user.Github
	user.lastUpdated = time.Now()

	for _, bot := range user.Bots {
		bot.Discord = "https://discord.gg/" + bot.Discord
	}

	k.CachedData.Users[id] = user
	return
}
