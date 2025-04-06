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

const (
	BotOk       BotState    = "ok"
	BotReported BotState    = "reported"
	BotBlocked  BotState    = "blocked"
	BotPrivate  BotState    = "private"
	BotArchived ServerState = "archived"
)

const (
	BotOnline    BotStatus = "online"
	BotIdle      BotStatus = "idle"
	BotDnD       BotStatus = "dnd"
	BotStreaming BotStatus = "streaming"
	BotOffline   BotStatus = "offline"
)
