package krlistsgo

import "encoding/json"

// UserFlags는 유저의 플래그 타입입니다.
type UserFlags int

// User는 한디리 API에서 반환된 유저 데이터입니다.
type User struct {
	ID       string         `json:"id"`
	Username string         `json:"username"`
	Tag      string         `json:"tag"`
	Github   string         `json:"github"`
	Flags    UserFlags      `json:"flags"`
	Bots     []BotInUser    `json:"bots"`
	Servers  []ServerInUser `json:"servers"`
}

// UserInBot은 Bot 구조체안에서의 User 구조체입니다.
type UserInBot struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Tag      string    `json:"tag"`
	Github   string    `json:"github"`
	Flags    UserFlags `json:"flags"`
	Bots     []string  `json:"bots"`
	Servers  []string  `json:"servers"`
}

// UserInServer는 Server 구조체안에서의 User 구조체입니다.
type UserInServer struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Tag      string    `json:"tag"`
	Github   string    `json:"github"`
	Flags    UserFlags `json:"flags"`
	Bots     []string  `json:"bots"`
	Servers  []string  `json:"servers"`
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
func (k *KrLists) User(id string) (user *User, err error) {
	resp, err := get(k.Client, "/users/"+id, []map[string]string{})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &user)
	return
}
