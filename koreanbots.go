package krlistsgo

import (
	"encoding/json"
	"net/http"
)

// 한디리 Go SDK의 구조체입니다.
type Koreanbots struct {
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

// New 메소드는 새로운 Koreanbots 구조체를 생성합니다은.
func New(token, clientID string) *Koreanbots {
	k := &Koreanbots{
		Token:    token,
		Client:   &http.Client{},
		ClientID: clientID,
	}
	return k
}
