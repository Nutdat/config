package core

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"
)

var (
	loadCache = make(map[string]interface{})
	loadMu    sync.RWMutex
)

func Load[T any](name string, out *T) error {
	// Prüfen, ob schon geladen
	loadMu.RLock()
	cached, ok := loadCache[name]
	loadMu.RUnlock()
	if ok {
		// Pointer aus dem Cache auf out kopieren
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(cached).Elem())
		return nil
	}

	// Standard Config prüfen
	mu.RLock()
	def, ok := registry[name]
	mu.RUnlock()
	if !ok {
		return fmt.Errorf("no default config registered under name %s", name)
	}

	path := fmt.Sprintf("./.Nutdat/config/%s.json", name)

	var cfg T
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := createConfigFile(name, def); err != nil {
			return err
		}
		cfg = def.(T)
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
	}

	// Cache setzen (Pointer speichern)
	loadMu.Lock()
	loadCache[name] = &cfg
	loadMu.Unlock()

	// out zeigt auf dieselbe Instanz
	reflect.ValueOf(out).Elem().Set(reflect.ValueOf(&cfg).Elem())
	return nil
}
