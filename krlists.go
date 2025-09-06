package krlistsgo

import (
	"encoding/json"
	"net/http"
)

// KrLists는 한디리 Go SDK의 구조체입니다.
type KrLists struct {
	Client         *http.Client
	BotIdentify    *Identify
	ServerIdentify *Identify
	CachedData     CachedData

	// 기본 캐시 간격은 1 시간 (3600 초)입니다.
	// 캐시 간격은 SetCache로 설정할 수 있습니다.
	CacheInterval int

	// Deprecated: BotIdentify 또는 ServerIdentify 구조체를 이용해주시길 바랍니다.
	Token string
	// Deprecated: BotIdentify 또는 ServerIdentify 구조체를 이용해주시길 바랍니다.
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
		CachedData: CachedData{
			Bots:    map[string]*Bot[User[string, string]]{},
			Servers: map[string]*Server[User[string, string]]{},
			Users:   map[string]*User[Bot[string], Server[string]]{},
		},
		CacheInterval: 3600,

		// 아래 값들은 SetBotIdentify를 써야 채워짐.
		Token:    "",
		ClientID: "",
	}
	return k
}

// SetBotIdentify는 당신의 봇의 자격 증명을 추가합니다.
func (k *KrLists) SetBotIdentify(token, id string) *KrLists {
	k.BotIdentify = &Identify{
		Token: token,
		ID:    id,
	}

	k.ClientID = id
	k.Token = token
	return k
}

// SetServerIdentify는 당신의 서버의 자격 증명을 추가합니다.
func (k *KrLists) SetServerIdentify(token, id string) *KrLists {
	k.ServerIdentify = &Identify{
		Token: token,
		ID:    id,
	}
	return k
}

func (k *KrLists) getAnyToken() (string, error) {
	if k.BotIdentify == nil {
		if k.ServerIdentify == nil {
			return "", ErrAnyTokenIsNil
		} else {
			return k.ServerIdentify.Token, nil
		}
	} else {
		return k.BotIdentify.Token, nil
	}
}
