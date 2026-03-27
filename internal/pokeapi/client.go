package pokeapi

import(
	"net/http"
	"time"
	"github.com/footlongdingledong/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient 	http.Client
	Cache		pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	client := Client{
		httpClient:	http.Client{
			Timeout: timeout,
		},
		Cache: pokecache.NewCache(cacheInterval),
	}
	return client
}
