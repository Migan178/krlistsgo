package koreanbotsgo

import "net/http"

const API_URL = "https://koreanbots.dev/api/v2"

type Koreanbots struct {
	Token  string
	Client *http.Client
}

func New(token, clientId string) *Koreanbots {
	k := &Koreanbots{
		Token:  token,
		Client: &http.Client{},
	}
	return k
}
