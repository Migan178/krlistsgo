package krlistsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 한디리의 API 주소입니다.
const API_URL = "https://koreanbots.dev/api"

// 한디리의 API 버전입니다.
const API_VERSION = "v2"

func get(client *http.Client, url string, headers []map[string]string) (data *ResponseBody, err error) {
	req, err := http.NewRequest(http.MethodGet, API_URL+"/"+API_VERSION+url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

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
		err = fmt.Errorf("http status code: %d, message: %s", data.Code, string(data.Message))
		return
	}
	return
}

func post(client *http.Client, url string, body any, headers []map[string]string) (data *ResponseBody, err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(http.MethodPost, API_URL+"/"+API_VERSION+url, bodyBuffer)
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return
	}

	if data.Code != 200 {
		err = fmt.Errorf("http Status Code: %d, message: %s", data.Code, string(data.Message))
		return
	}
	return
}
