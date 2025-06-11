package pokeapi

import (
	"net/http"
	"time"

	"github.com/db-cooper7/bootdev-guided-projects/go-pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timemout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timemout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
