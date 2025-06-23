package cfg

import "github.com/WebPirat/nutdat-config/package/core"

// Register exposes the core register functionality
func Register[T any](name string, defaultCfg T) error {
	return core.Register(name, defaultCfg)
}

// Load exposes the core load functionality
func Load[T any](name string, out *T) error {
	return core.Load(name, out)
}
