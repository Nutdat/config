package main

import (
	"fmt"

	"github.com/Nutdat/config"
)

var loaded *AppConfig

// Nested LoggerConfig
type LoggerConfig struct {
	Level string `json:"level"`
	Path  string `json:"path"`
}

// Top-level AppConfig mit embedded LoggerConfig
type AppConfig struct {
	AppName string       `json:"app_name"`
	Debug   bool         `json:"debug"`
	Logger  LoggerConfig `json:"logger"`
}

func init() {
	defaultConfig := AppConfig{
		AppName: "MyCoolApp",
		Debug:   true,
		Logger: LoggerConfig{
			Level: "info",
			Path:  "./logs/app.log",
		},
	}

	if err := cfg.Register("app", defaultConfig); err != nil {
		panic(err)
	}

	if err := cfg.Load("app", &loaded); err != nil {
		panic(err)
	}
}
func LoadConfig() *AppConfig {
	fmt.Printf("cfg ptr: %p\n", loaded)
	return loaded
}
