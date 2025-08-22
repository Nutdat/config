package core

import (
	"encoding/json"
	"fmt"
	"os"
)

func createConfigFile(name string, cfg any) error {
	path := fmt.Sprintf("./.Nutdat/config/%s.json", name)

	if err := os.MkdirAll("./.Nutdat/config", 0755); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		if err := enc.Encode(cfg); err != nil {
			return err
		}
	}

	return nil
}
