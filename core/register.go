package core

import "sync"

var (
    registry = make(map[string]any)
    mu       sync.RWMutex
)

func Register[T any](name string, defaultCfg T) error {
    mu.Lock()
    defer mu.Unlock()

    registry[name] = defaultCfg
    return createConfigFile(name, defaultCfg)
}

