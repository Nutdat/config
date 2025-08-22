package core

import (
	"fmt"
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
		return fmt.Errorf("config %q already registered", name)
	}

	registry[name] = defaultCfg
	return createConfigFile(name, defaultCfg)
}
