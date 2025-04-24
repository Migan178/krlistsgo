package krlistsgo

import "fmt"

// 한디리의 API 주소입니다.
const API_URL = "https://koreanbots.dev/api"

// 한디리의 API 버전입니다.
const API_VERSION = 2

// 알려진 한디리의 엔드 포인트들입니다. (위젯 제외)
var (
	EndpointBots             = func(id string) string { return fmt.Sprintf("%s/v%d/bots/%s", API_URL, API_VERSION, id) }
	EndpointBotUpdateServers = func(id string) string { return fmt.Sprintf("%s/stats", id) }
	EndpointBotVote          = func(botID, userID string) string {
		return fmt.Sprintf("%s/vote?userID=%s", EndpointBots(botID), userID)
	}

	EndpointServers      = func(id string) string { return fmt.Sprintf("%s/v%d/servers/%s", API_URL, API_VERSION, id) }
	EndpointServerOwners = func(id string) string { return fmt.Sprintf("%s/owners", EndpointServers(id)) }
	EndpointServerVote   = func(serverID, userID string) string {
		return fmt.Sprintf("%s/vote?userID=%s", EndpointServers(serverID), userID)
	}

	EndpointNewBots    = func(page int) string { return fmt.Sprintf("%s/v%d/lists/bots/new?page=%d", API_URL, API_VERSION, page) }
	EndpointSearchBots = func(query string, page int) string {
		return fmt.Sprintf("%s/v%d/search/bots?query=%s&page=%d", API_URL, API_VERSION, query, page)
	}
	EndpointBotRanking = func(page int) string {
		return fmt.Sprintf("%s/v%d/lists/bots/vote?page=%d", API_URL, API_VERSION, page)
	}

	EndpointUsers = func(id string) string { return fmt.Sprintf("%s/v%d/users/%s", API_URL, API_VERSION, id) }

	EndpointSearchServers = func(query string, page int) string {
		return fmt.Sprintf("%s/v%d/search/servers?query=%s&page=%d", API_URL, API_VERSION, query, page)
	}
)
