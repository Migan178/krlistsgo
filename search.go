package krlistsgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// SearchResult는 검색한 결과 값입니다.
type SearchResult[T Server[User[string, string]] | Bot[string]] struct {
	Data       []T          `json:"Data"`
	Current    int          `json:"currentPage"`
	Total      int          `json:"totalPage"`
	client     *http.Client `json:"-"`
	searchType string       `json:"-"`
	query      string       `json:"-"`
}

func search[T Server[User[string, string]] | Bot[string]](c *http.Client, searchType string, query string, page int) (result *SearchResult[T], err error) {
	query = url.QueryEscape(query)
	resp, err := get(c, fmt.Sprintf("/search/%s?query=%s&page=%d", searchType, query, page), []map[string]string{})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &result)
	result.client = c
	result.searchType = searchType
	result.query = query
	return
}

// SearchBots는 해당 검색어로 봇을 검색합니다.
func (k *KrLists) SearchBots(query string) (*SearchResult[Bot[string]], error) {
	return search[Bot[string]](k.Client, "bots", query, 1)
}

// SearchServers는 해당 검색어로 서버를 검색합니다.
func (k *KrLists) SearchServers(query string) (*SearchResult[Server[User[string, string]]], error) {
	return search[Server[User[string, string]]](k.Client, "servers", query, 1)
}

// Next는 다음 페이지로 넘깁니다.
func (s *SearchResult[T]) Next() (*SearchResult[T], error) {
	if s.Current >= s.Total {
		return nil, SearchLastPageErr
	}
	return search[T](s.client, s.searchType, s.query, s.Current+1)
}

// Previous는 이전 페이지로 넘깁니다.
func (s *SearchResult[T]) Previous() (*SearchResult[T], error) {
	if s.Current <= 1 {
		return nil, SearchFirstPageErr
	}
	return search[T](s.client, s.searchType, s.query, s.Current-1)
}

// Select는 특정 페이지로 바로 이동합니다.
func (s *SearchResult[T]) Select(page int) (*SearchResult[T], error) {
	if page < 0 {
		return nil, SearchPositiveNumberErr
	}

	if page == s.Current {
		return s, nil
	}
	return search[T](s.client, s.searchType, s.query, page)
}
