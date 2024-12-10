package cache

import "time"

const cacheEntryInvalidationHours = 12

var internalCache map[string]PriceCacheEntry

type PriceCacheEntry struct {
	price       float64
	lastUpdated time.Time
}

func init() {
	internalCache = map[string]PriceCacheEntry{}
}

func Get(key string) (float64, bool) {
	if entry, ok := internalCache[key]; ok {
		if entry.Expired() {
			return entry.price, true
		}
		// it os not, clear the value
		delete(internalCache, key)
	}
	return 0.0, false
}

func Store(key string, price float64) {
	entry := PriceCacheEntry{
		price:       price,
		lastUpdated: time.Now(),
	}
	internalCache[key] = entry
}

func (e PriceCacheEntry) Expired() bool {
	now := time.Now()
	if now.Sub(e.lastUpdated).Hours() < cacheEntryInvalidationHours {
		return true
	}
	return false
}
