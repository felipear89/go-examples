package cache

import . "github.com/patrickmn/go-cache"

func WithCache[V any](cache *Cache, k string, f func() (V, error)) (V, error) {
	if v, ok := cache.Get(k); ok {
		return v.(V), nil
	}
	v, err := f()
	if err != nil {
		var zero V
		return zero, err
	}
	cache.Set(k, v, DefaultExpiration)
	return v, nil
}
