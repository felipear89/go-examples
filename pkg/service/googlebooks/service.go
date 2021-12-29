package googlebooks

import (
	"github.com/felipear89/go-examples/pkg/cache"
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type SearchBooks func(query string) (*Response, error)

func NewSearchBooks(c RequestSearchBooks) SearchBooks {
	memoryCache := gocache.New(5*time.Minute, 10*time.Minute)
	return func(query string) (*Response, error) {
		return cache.WithCache(memoryCache, query, searchBooks(c, query))
	}
}

func searchBooks(requestSearchBooks RequestSearchBooks, query string) func() (*Response, error) {
	return func() (*Response, error) {
		return requestSearchBooks(query)
	}
}
