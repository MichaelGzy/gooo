package client

import (
	cache "github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}
