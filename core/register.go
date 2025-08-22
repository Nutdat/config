package core

import (
	"sync"
)

var (
	registry = make(map[string]any)
	mu       sync.RWMutex
)

func Register[T any](name string, defaultCfg T) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := registry[name]; exists {
		// Schon registriert → Fehler zurückgeben
		return nil
	}

	registry[name] = defaultCfg
	return createConfigFile(name, defaultCfg)
}
