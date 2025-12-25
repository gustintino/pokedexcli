package pokeapi

import (
	"net/http"
	"time"

	"github.com/gustintino/pokedexcli/internal/pokecache"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	Client http.Client
	cache  pokecache.Cache
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		Client: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(10 * time.Second),
	}
}
