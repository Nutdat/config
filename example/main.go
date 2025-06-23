package main

import (
	"fmt"

	"github.com/Nutdat/config"
)

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

func main() {
	defaultConfig := AppConfig{
		AppName: "MyCoolApp",
		Debug:   true,
		Logger: LoggerConfig{
			Level: "info",
			Path:  "./logs/app.log",
		},
	}

	// Register mit verschachtelter Konfiguration
	err := cfg.Register("app", defaultConfig)
	if err != nil {
		panic(err)
	}

	var loaded AppConfig
	if err := cfg.Load("app", &loaded); err != nil {
		panic(err)
	}

	fmt.Println("== Loaded Config ==")
	fmt.Printf("App Name: %s\n", loaded.AppName)
	fmt.Printf("Debug:    %v\n", loaded.Debug)
	fmt.Printf("Logger:\n  Level: %s\n  Path:  %s\n",
		loaded.Logger.Level, loaded.Logger.Path)
}
