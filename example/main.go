package main

import "fmt"

func main() {
	loaded := LoadConfig()
	loaded2 := LoadConfig()
	fmt.Printf("cfg1 ptr: %p\n", loaded)
	fmt.Printf("cfg2 ptr: %p\n", loaded2)
	fmt.Println("== Loaded Config ==")
	fmt.Printf("App Name: %s\n", loaded.AppName)
	fmt.Printf("Debug:    %v\n", loaded.Debug)
	fmt.Printf("Logger:\n  Level: %s\n  Path:  %s\n",
		loaded.Logger.Level, loaded.Logger.Path)
}
