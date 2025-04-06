package koreanbotsgo

type UserFlags int

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Tag      string    `json:"tag"`
	Github   string    `json:"github"`
	Flags    UserFlags `json:"flags"`
	Bots     []*Bot    `json:"bots"`
	Servers  []*Server `json:"servers"`
}

const (
	UserNone          UserFlags = 0 << 0
	UserAdministrator UserFlags = 1 << 0
	UserBugHunter     UserFlags = 1 << 1
	UserReviewer      UserFlags = 1 << 2
	UserPremium       UserFlags = 1 << 3
)
