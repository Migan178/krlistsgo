package krlistsgo

import "encoding/json"

// UserFlags는 유저의 플래그 타입입니다.
type UserFlags int

// User는 한디리 API에서 반환된 유저 데이터입니다.
type User[T, V any] struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Tag      string    `json:"tag"`
	Github   string    `json:"github"`
	Flags    UserFlags `json:"flags"`
	Bots     []T       `json:"bots"`
	Servers  []V       `json:"servers"`
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
	resp, err := get(k.Client, "/users/"+id, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &user)

	user.Github = "https://github.com/" + user.Github

	for _, bot := range user.Bots {
		bot.Discord = "https://discord.gg/" + bot.Discord
	}
	return
}
