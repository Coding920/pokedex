package pokeapi

import (
	"github.com/Coding920/pokedex/internal/pokecache"
	"net/http"
	"time"
)

const (
	baseApi = "https://pokeapi.co/api/v2/"
)

type PokeClient struct {
	client http.Client
	cache  pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) PokeClient {
	return PokeClient{
		client: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
