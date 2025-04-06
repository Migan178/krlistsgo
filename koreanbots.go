package koreanbotsgo

import (
	"encoding/json"
	"net/http"
)

// 한디리 Go SDK의 구조체입니다.
type Koreanbots struct {
	Token  string
	Client *http.Client
}

// 한디리 API에서 반환되는 데이터입니다.
type ResponseBody struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message json.RawMessage `json:"message"`
	Version int             `json:"version"`
}

// 한디리의 API 주소입니다.
const API_URL = "https://koreanbots.dev/api/v2"

// New 메소드는 새로운 Koreanbots 구조체를 생성합니다은.
func New(token, clientId string) *Koreanbots {
	k := &Koreanbots{
		Token:  token,
		Client: &http.Client{},
	}
	return k
}
