package krlistsgo

type CachedData struct {
	Bots    map[string]*Bot[User[string, string]]
	Servers map[string]*Server[User[string, string]]
	Users   map[string]*User[Bot[string], Server[string]]
}

// SetCache는 krlistsgo의 캐시 간격을 조절합니다.
// 값을 0보다 작게 설정하면 캐시 기능이 꺼집니다.
// krlistsgo의 기본 캐시 간격은 1 시간 (3600 초)입니다.
// interval값은 int형으로 초 단위로 설정해야합니다. (예시: 1 (시간) -> 3600 (초))
func (k *KrLists) SetCache(interval int) *KrLists {
	if interval < 1 {
		k.CacheInterval = 0
	}

	k.CacheInterval = interval
	return k
}
