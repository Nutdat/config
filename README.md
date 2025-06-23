# 🧩 Nut Config 

A flexible configuration management utility for Go modules.  
This package allows you to **register**, **persist**, and **load** configuration structs using simple JSON files.  
Each module can define and own its config without hardcoding logic for file management.

---

## 📦 Features

- 📁 Configs stored as JSON files in `./data/config/`
- 🧩 Modular: Each package registers its own default config
- ⚙️ Auto-creates config files on first use
- 🔄 Reload configs easily at runtime
- 🧪 Unit-tested and production-ready

---

## 🚀 Quickstart

### 1. Define your config struct

```go
type LoggerConfig struct {
    Level string `json:"level"`
    Path  string `json:"path"`
}
```

### 2. Register the default config

```go
cfg.Register("logger", LoggerConfig{
    Level: "info",
    Path:  "./logs/app.log",
})
```

### 3. Load the config 

```go
var c LoggerConfig
cfg.Load("logger", &c)
fmt.Println(c.Level, c.Path)
```


### Running the tests

```
go test ./package/core
```

### Struct

```
├── cfg.go            <- public API (Register, Load)
├── /core             <- create.go, load.go, register.go
├── /example          <- demo with main.go
```
