package provider

import "sync"

var (
	registry = make(map[string]Provider)
	mu       sync.RWMutex
)

func RegisterProviders(name string, provider Provider) {
	mu.Lock()
	defer mu.Unlock()
	registry[name] = provider
}

func GetProvider(name string) (Provider, bool) {
	mu.RLock()
	defer mu.RUnlock()
	provider, ok := registry[name]
	return provider, ok
}
