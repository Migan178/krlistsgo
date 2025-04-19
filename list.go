package krlistsgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// ListType은 리스트의 정보 타입입니다.
type ListType string

// PageResult는 리스트의 결과 값입니다.
type PageResult[T Server[User[string, string]] | Bot[string]] struct {
	Type       ListType     `json:"type"`
	Data       []T          `json:"Data"`
	Current    int          `json:"currentPage"`
	Total      int          `json:"totalPage"`
	client     *http.Client `json:"-"`
	searchType string       `json:"-"`
	query      string       `json:"-"`
}

// 리스트의 타입입니다.
const (
	ListNew    ListType = "NEW"
	ListSearch ListType = "SEARCH"
	ListVote   ListType = "VOTE"
)

func search[T Server[User[string, string]] | Bot[string]](c *http.Client, searchType string, query string, page int) (result *PageResult[T], err error) {
	if page <= 0 {
		return nil, ListPositiveNumberErr
	}

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

func new[T Server[User[string, string]] | Bot[string]](c *http.Client, searchType string, page int) (result *PageResult[T], err error) {
	if page <= 0 {
		return nil, ListPositiveNumberErr
	}

	resp, err := get(c, fmt.Sprintf("/list/%s/new?page=%d", searchType, page), []map[string]string{})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &result)
	result.client = c
	result.searchType = searchType
	return
}

func votes[T Server[User[string, string]] | Bot[string]](c *http.Client, searchType string, page int) (result *PageResult[T], err error) {
	if page <= 0 {
		return nil, ListPositiveNumberErr
	}

	resp, err := get(c, fmt.Sprintf("/list/%s/votes?page=%d", searchType, page), []map[string]string{})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp.Data, &result)
	result.client = c
	result.searchType = searchType
	return
}

// SearchBots는 해당 검색어로 봇을 검색합니다.
func (k *KrLists) SearchBots(query string, page int) (*PageResult[Bot[string]], error) {
	return search[Bot[string]](k.Client, "bots", query, page)
}

// SearchServers는 해당 검색어로 서버를 검색합니다.
func (k *KrLists) SearchServers(query string, page int) (*PageResult[Server[User[string, string]]], error) {
	return search[Server[User[string, string]]](k.Client, "servers", query, page)
}

func (k *KrLists) NewBots(page int) (*PageResult[Bot[string]], error) {
	return new[Bot[string]](k.Client, "bots", page)
}

func (k *KrLists) NewServers(page int) (*PageResult[Server[User[string, string]]], error) {
	return new[Server[User[string, string]]](k.Client, "servers", page)
}

func (k *KrLists) BotsVoteRanking(page int) (*PageResult[Bot[string]], error) {
	return votes[Bot[string]](k.Client, "bots", page)
}

func (k *KrLists) ServersVoteRanking(page int) (*PageResult[Server[User[string, string]]], error) {
	return votes[Server[User[string, string]]](k.Client, "servers", page)
}

// Next는 다음 페이지로 넘깁니다.
func (p *PageResult[T]) Next() (*PageResult[T], error) {
	if p.Current >= p.Total {
		return nil, ListLastPageErr
	}

	switch p.Type {
	case ListNew:
		return new[T](p.client, p.searchType, p.Current+1)
	case ListSearch:
		return search[T](p.client, p.searchType, p.query, p.Current+1)
	case ListVote:
		return votes[T](p.client, p.searchType, p.Current+1)
	}
	return nil, nil
}

// Previous는 이전 페이지로 넘깁니다.
func (p *PageResult[T]) Previous() (*PageResult[T], error) {
	if p.Current <= 1 {
		return nil, ListFirstPageErr
	}

	switch p.Type {
	case ListNew:
		return new[T](p.client, p.searchType, p.Current+1)
	case ListSearch:
		return search[T](p.client, p.searchType, p.query, p.Current+1)
	case ListVote:
		return votes[T](p.client, p.searchType, p.Current+1)
	}
	return nil, nil
}

// Select는 특정 페이지로 바로 이동합니다.
func (p *PageResult[T]) Select(page int) (*PageResult[T], error) {
	if page == p.Current {
		return p, nil
	}

	switch p.Type {
	case ListNew:
		return new[T](p.client, p.searchType, page)
	case ListSearch:
		return search[T](p.client, p.searchType, p.query, page)
	case ListVote:
		return votes[T](p.client, p.searchType, page)
	}
	return nil, nil
}

// First는 가장 첫 페이지로 이동합니다.
func (p *PageResult[T]) First() (*PageResult[T], error) {
	if p.Current == 1 {
		return p, nil
	}

	switch p.Type {
	case ListNew:
		return new[T](p.client, p.searchType, 1)
	case ListSearch:
		return search[T](p.client, p.searchType, p.query, 1)
	case ListVote:
		return votes[T](p.client, p.searchType, 1)
	}
	return nil, nil
}

// Last는 가장 마지막 페이지로 이동합니다.
func (p *PageResult[T]) Last() (*PageResult[T], error) {
	if p.Current == p.Total {
		return p, nil
	}

	switch p.Type {
	case ListNew:
		return new[T](p.client, p.searchType, p.Total)
	case ListSearch:
		return search[T](p.client, p.searchType, p.query, p.Total)
	case ListVote:
		return votes[T](p.client, p.searchType, p.Total)
	}
	return nil, nil
}
