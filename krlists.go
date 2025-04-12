package krlistsgo

import (
	"encoding/json"
	"net/http"
)

// 한디리 Go SDK의 구조체입니다.
type KrLists struct {
	Token    string
	Client   *http.Client
	ClientID string
}

// 한디리 API에서 반환되는 데이터입니다.
type ResponseBody struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message json.RawMessage `json:"message"`
	Version int             `json:"version"`
}

// CheckVote는 해당 봇이나 서버의 투표 여부를 확인할 때 반환되는 데이터입니다.
type CheckVote struct {
	Voted    bool `json:"voted"`
	LastVote int  `json:"lastvote"`
}

// New 메소드는 새로운 Krlists 구조체를 생성합니다.
func New(token, clientID string) *KrLists {
	k := &KrLists{
		Token:    token,
		Client:   &http.Client{},
		ClientID: clientID,
	}
	return k
}
