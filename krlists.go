package krlistsgo

import (
	"encoding/json"
	"net/http"
)

// 한디리 Go SDK의 구조체입니다.
type KrLists struct {
	Client         *http.Client
	BotIdentify    *Identify
	ServerIdentify *Identify

	// Deprecated: BotIdentify 또는 ServerIdentify 구조체를 이용해주세요.
	Token string
	// Deprecated: BotIdentify 또는 ServerIdentify 구조체를 이용해주세요.
	ClientID string
}

// 한디리의 API 자격 증명입니다.
type Identify struct {
	ID    string
	Token string
}

// 한디리 API에서 반환되는 데이터입니다.
type ResponseBody struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message json.RawMessage `json:"message"`
	Version int             `json:"version"`
}

// New 메소드는 새로운 KrLists 구조체를 생성합니다.
func New() *KrLists {
	k := &KrLists{
		Client: &http.Client{},

		// 아래 값들은 SetBotIdentify를 써야 채워짐.
		Token:    "",
		ClientID: "",
	}
	return k
}

func (k *KrLists) SetBotIdentify(token, id string) *KrLists {
	k.BotIdentify = &Identify{
		Token: token,
		ID:    id,
	}

	k.ClientID = id
	k.Token = token
	return k
}

func (k *KrLists) SetServerIdentify(token, id string) *KrLists {
	k.ServerIdentify = &Identify{
		Token: token,
		ID:    id,
	}
	return k
}
