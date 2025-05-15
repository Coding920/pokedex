package pokeapi

import (
	"net/http"
	"time"
)

const (
	baseApi = "https://pokeapi.co/api/v2/"
)

type PokeClient struct {
	client http.Client
}

func NewClient(timeout time.Duration) PokeClient {
	return PokeClient{
		client: http.Client{
			Timeout: timeout,
		},
	}
}

type LocationsJson struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}
