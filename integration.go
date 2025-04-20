package krlistsgo

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// NewKrListsWithDiscordGo 메소드는 discordgo를 사용해 자동 서버 수 업데이트를 할 수 있는 새로운 KrLists 구조체를 생성합니다.
func NewKrListsWithDiscordGo(s *discordgo.Session, token string, automaticUpdateServers bool) *KrLists {
	k := New().SetBotIdentify(token, s.State.User.ID)

	if automaticUpdateServers {
		ticker := time.NewTicker(time.Minute * 10)
		for range ticker.C {
			k.UpdateServers(len(s.State.Guilds), s.ShardCount)
		}
	}
	return k
}
