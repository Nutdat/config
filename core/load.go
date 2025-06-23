package core

import (
    "encoding/json"
    "fmt"
    "os"
    "reflect"
)

func Load[T any](name string, out *T) error {
    mu.RLock()
    def, ok := registry[name]
    mu.RUnlock()

    if !ok {
        return fmt.Errorf("no default config registered under name %s", name)
    }

    path := fmt.Sprintf("./data/config/%s.json", name)

    if _, err := os.Stat(path); os.IsNotExist(err) {
        if err := createConfigFile(name, def); err != nil {
            return err
        }
        reflect.ValueOf(out).Elem().Set(reflect.ValueOf(def))
        return nil
    }

    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }

    return json.Unmarshal(data, out)
}

