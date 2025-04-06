package koreanbotsgo

type BotFlags int
type BotLib string
type BotState string
type BotStatus string
type BotCategory string

type Bot struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Tag      string      `json:"tag"`
	Avatar   string      `json:"avatar"`
	Owners   []User      `json:"owners"`
	Flags    BotFlags    `json:"flags"`
	Lib      BotLib      `json:"lib"`
	Prefix   string      `json:"prefix"`
	Votes    int         `json:"votes"`
	Servers  int         `json:"servers"`
	Shards   int         `json:"shards"`
	Intro    string      `json:"intro"`
	Desc     string      `json:"desc"`
	Web      string      `json:"web"`
	Git      string      `json:"git"`
	Url      string      `json:"url"`
	Discord  string      `json:"discord"`
	Category BotCategory `json:"Category"`
	Vanity   string      `json:"vanity"`
	Bg       string      `json:"bg"`
	Banner   string      `json:"banner"`
	Status   BotStatus   `json:"status"`
	State    BotState    `json:"state"`
}

const (
	BotNone            BotFlags = 0 << 0
	BotOfficial        BotFlags = 1 << 0
	BotKrBotsVerified  BotFlags = 1 << 2
	BotPartner         BotFlags = 1 << 3
	BotDiscordVerified BotFlags = 1 << 4
	BotPremium         BotFlags = 1 << 5
	BotHackathon       BotFlags = 1 << 6
)

const (
	DiscordJs     BotLib = "discord.js"
	Eris          BotLib = "Eris"
	DiscordPy     BotLib = "discord.py"
	DiscordCr     BotLib = "discordcr"
	Nyxx          BotLib = "Nyxx"
	DiscordDotNet BotLib = "discord.net"
	DSharpPlus    BotLib = "DSharpPlus"
	Nostrum       BotLib = "Nostrum"
	Coxir         BotLib = "coxir"
	DiscordGo     BotLib = "DiscordGo"
	Discord4J     BotLib = "Discord4J"
	Javacord      BotLib = "Javacord"
	JDA           BotLib = "JDA"
	Discordia     BotLib = "Discordia"
	RestCord      BotLib = "RestCord"
	Yasmin        BotLib = "Yasmin"
	Disco         BotLib = "disco"
	Discordrb     BotLib = "discordrb"
	Serenity      BotLib = "serenity"
	SwiftDiscord  BotLib = "SwiftDiscord"
	Sword         BotLib = "Sword"
	Other         BotLib = "기타"
	Private       BotLib = "비공개"
)
