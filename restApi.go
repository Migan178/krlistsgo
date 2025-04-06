package koreanbotsgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 한디리의 API 주소입니다.
const API_URL = "https://koreanbots.dev/api/v2"

func get(client *http.Client, url string, headers []map[string]string) (data *ResponseBody, err error) {
	req, err := http.NewRequest(http.MethodGet, API_URL+url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	if len(headers) > 0 {
		for _, header := range headers {
			for key, value := range header {
				req.Header.Add(key, value)
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	if data.Code != 200 {
		err = errors.New(fmt.Sprintf("Http Status Code: %d, Message: %s", data.Code, string(data.Message)))
		return
	}

	return
}
