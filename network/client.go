package network

import (
	"net/http"
	"time"

	"github.com/Hessam-Emami/pokedex/cache"
)

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
