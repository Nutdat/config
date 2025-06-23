package core

import (
    "encoding/json"
    "os"
    "path/filepath"
    "testing"
)

type dummyConfig struct {
    Foo string `json:"foo"`
}

func TestCreateConfigFile_CreatesJSONFile(t *testing.T) {
    name := "test_create"
    path := filepath.Join("./data/config", name+".json")
    _ = os.Remove(path) 

    cfg := dummyConfig{Foo: "bar"}

    err := createConfigFile(name, cfg)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    defer os.Remove(path) 

    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("failed to read file: %v", err)
    }

    var decoded dummyConfig
    if err := json.Unmarshal(data, &decoded); err != nil {
        t.Fatalf("json parse error: %v", err)
    }

    if decoded.Foo != "bar" {
        t.Errorf("expected Foo = bar, got %s", decoded.Foo)
    }
}

