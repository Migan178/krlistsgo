package krlistsgo

import (
	"encoding/json"
	"net/http"
)

// CheckVote는 해당 봇이나 서버의 투표 여부를 확인할 때 반환되는 데이터입니다.
type CheckVote struct {
	Voted    bool `json:"voted"`
	LastVote int  `json:"lastvote"`
}

func checkVote(c *http.Client, i *Identify, voteType, userID string) (data *CheckVote, err error) {
	if i == nil {
		if voteType == "bots" {
			return nil, BotIdentifyIsNil
		}

		return nil, ServerIdentifyIsNil
	}

	resp, err := get(c, "/"+voteType+"/"+i.ID+"/vote?userID="+userID, &map[string]string{
		"Authorization": i.Token,
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &data)
	return
}

// CheckBotVote는 userID가 당신의 봇에 투표했는지를 확인합니다.
func (k *KrLists) CheckBotVote(userID string) (*CheckVote, error) {
	return checkVote(k.Client, k.BotIdentify, "bots", userID)
}

// CheckVote는 userID가 해당 봇에 투표했는지를 확인합니다.
// token 인수는 해당 봇이 당신의 봇인지를 모르기 때문에 필요합니다.
func (b *Bot[T]) CheckVote(token, userID string) (*CheckVote, error) {
	return checkVote(b.client, &Identify{b.ID, token}, "bots", userID)
}

// CheckServerVote는 userID가 당신의 서버에 투표했는지를 확인합니다.
func (k *KrLists) CheckServerVote(userID string) (*CheckVote, error) {
	return checkVote(k.Client, k.ServerIdentify, "servers", userID)
}

// CheckVote는 userID가 해당 서버에 투표했는지를 확인합니다.
// token 인수는 해당 봇이 당신의 서버인지를 모르기 때문에 필요합니다.
func (s *Server[T]) CheckVote(token, userID string) (*CheckVote, error) {
	return checkVote(s.client, &Identify{s.ID, token}, "servers", userID)
}
