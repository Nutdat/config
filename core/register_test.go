package core

import (
    "encoding/json"
    "os"
    "path/filepath"
    "testing"
)

func TestRegister_StoresDefaultAndCreatesFile(t *testing.T) {
    name := "test_register"
    path := filepath.Join("./.Nutdat/config", name+".json")
    _ = os.Remove(path)

    type Config struct {
        Value int `json:"value"`
    }

    def := Config{Value: 123}
    err := Register(name, def)
    if err != nil {
        t.Fatalf("Register failed: %v", err)
    }

    defer os.Remove(path)

    // Check ob registriert
    mu.RLock()
    stored, ok := registry[name]
    mu.RUnlock()

    if !ok {
        t.Fatal("config was not store d in registry")
    }

    val, ok := stored.(Config)
    if !ok || val.Value != 123 {
        t.Errorf("unexpected registry value: %+v", stored)
    }

    // Check Datei
    raw, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("config file not created: %v", err)
    }

    var loaded Config
    _ = json.Unmarshal(raw, &loaded)
    if loaded.Value != 123 {
        t.Errorf("file content mismatch: %+v", loaded)
    }
}

