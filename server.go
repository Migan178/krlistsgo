package koreanbotsgo

type ServerFlags int
type ServerState string
type ServerCategory string

type Server struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Icon      string           `json:"icon"`
	Owner     User             `json:"owner"`
	Flags     ServerFlags      `json:"flags"`
	Votes     int              `json:"votes"`
	Members   int              `json:"members"`
	BoostTier int              `json:"boostTier"`
	Intro     string           `json:"intro"`
	Desc      string           `json:"desc"`
	Category  []ServerCategory `json:"category"`
	Invite    string           `json:"invite"`
	Emojis    []*ServerEmoji   `json:"emojis"`
	Vanity    string           `json:"vanity"`
	Bg        string           `json:"bg"`
	Banner    string           `json:"banner"`
	State     ServerState      `json:"state"`
}

type ServerEmoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

const (
	ServerNone            ServerFlags = 0 << 0
	ServerOfficial        ServerFlags = 1 << 0
	ServerKrBotsVerified  ServerFlags = 1 << 2
	ServerKrBotsPartner   ServerFlags = 1 << 3
	ServerDiscordVerified ServerFlags = 1 << 4
	ServerDiscordPartner  ServerFlags = 1 << 5
)

const (
	ServerOk          ServerState = "ok"
	ServerReported    ServerState = "reported"
	ServerBlocked     ServerState = "blocked"
	ServerUnReachable ServerState = "unreachable"
)

const (
	ServerCommunity ServerCategory = "커뮤니티"
	ServerITScience ServerCategory = "IT & 과학"
	ServerBot       ServerCategory = "봇"
	ServerAmity     ServerCategory = "친목"
	ServerMusic     ServerCategory = "음악"
	ServerEducation ServerCategory = "교육"
	Serverdate      ServerCategory = "연애"
	ServerGame      ServerCategory = "게임"
	ServerOverwatch ServerCategory = "오버워치"
	ServerLOL       ServerCategory = "리그 오브 레전드"
	ServerPUBG      ServerCategory = "배틀그라운드"
	ServerMinecraft ServerCategory = "마인크래프트"
)
