package core

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_ReadsExistingFile(t *testing.T) {
	name := "test_load_existing"
	path := filepath.Join("./.Nutdat/config", name+".json")
	_ = os.MkdirAll("./.Nutdat/config", 0755)
	_ = os.Remove(path)

	type Config struct {
		Mode string `json:"mode"`
	}

	expected := Config{Mode: "release"}

	// Manuell Datei schreiben
	f, _ := os.Create(path)
	_ = json.NewEncoder(f).Encode(expected)
	_ = f.Close()
	defer os.Remove(path)

	// Registry mit anderem Default
	_ = Register(name, Config{Mode: "debug"})

	var loaded Config
	err := Load(name, &loaded)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.Mode != "debug" {
		t.Errorf("expected mode=debug, got %s", loaded.Mode)
	}
}

func TestLoad_CreatesFileIfNotExists(t *testing.T) {
	name := "test_load_create"
	path := filepath.Join("./.Nutdat/config", name+".json")
	_ = os.Remove(path)

	type Config struct {
		Enabled bool `json:"enabled"`
	}

	def := Config{Enabled: true}
	_ = Register(name, def)

	var cfg Config
	err := Load(name, &cfg)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if !cfg.Enabled {
		t.Errorf("expected Enabled = true, got false")
	}

	_, err = os.Stat(path)
	if err != nil {
		t.Errorf("config file not created: %v", err)
	}

	_ = os.Remove(path)
}
