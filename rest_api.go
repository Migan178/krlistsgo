package krlistsgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get(client *http.Client, url string, header *map[string]string) (data *ResponseBody, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	if header != nil {
		for key, value := range *header {
			req.Header.Add(key, value)
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

func post(client *http.Client, url string, body any, header *map[string]string) (data *ResponseBody, err error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return
	}
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest(http.MethodPost, url, bodyBuffer)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")

	if header != nil {
		for key, value := range *header {
			req.Header.Add(key, value)
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
