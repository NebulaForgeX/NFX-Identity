package configx

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"nebulaid/pkgs/utils/file"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/confmap"
	envProvider "github.com/knadh/koanf/providers/env"
	fileProvider "github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Loader[T any] struct {
	cfg *T
	k   *koanf.Koanf
}

type options struct {
	explicitPath string
	envPrefix    string
	watch        bool
	defaults     map[string]any
}

type LoaderOption func(*options)

func WithPath(p string) LoaderOption      { return func(o *options) { o.explicitPath = p } }
func WithEnvPrefix(p string) LoaderOption { return func(o *options) { o.envPrefix = p } }
func WithDefaults(m map[string]any) LoaderOption {
	return func(o *options) { o.defaults = m }
}

func NewLoader[T any](ctx context.Context, opts ...LoaderOption) (*Loader[T], error) {
	op := options{envPrefix: "APP"}
	for _, f := range opts {
		f(&op)
	}

	k := koanf.New(".")

	if len(op.defaults) > 0 {
		k.Load(confmap.Provider(op.defaults, "."), nil)
	}

	// === TOML provider ===
	path, err := resolvePath(op.explicitPath)
	if err != nil {
		return nil, err
	}
	fProvider := fileProvider.Provider(path)
	if err := k.Load(fProvider, toml.Parser()); err != nil {
		return nil, fmt.Errorf("load toml config failed: %w", err)
	}

	// === ENV provider ===
	replacer := func(s string) string {
		return strings.NewReplacer("__", "-", "_", ".").Replace(s)
	}
	if err := k.Load(envProvider.Provider(op.envPrefix+"_", ".", replacer), nil); err != nil {
		return nil, fmt.Errorf("load env config failed: %w", err)
	}

	// === Unmarshal config ===
	var newCfg T
	if err := k.Unmarshal("", &newCfg); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %w", err)
	}

	return &Loader[T]{cfg: &newCfg, k: k}, nil
}

func (l *Loader[T]) Config() *T {
	return l.cfg
}

func (l *Loader[T]) PrintAll() {
	l.k.Print()
}

func (l *Loader[T]) PrintWithPrefix(prefix string) {
	keys := l.k.Keys()
	for _, key := range keys {
		if strings.HasPrefix(key, prefix) {
			fmt.Printf("%s -> %v\n", key, l.k.Get(key))
		}
	}
}

func resolvePath(explicit string) (string, error) {
	if explicit != "" && file.Exists(explicit) {
		return explicit, nil
	}

	// --flag
	if f := flag.Lookup("config"); f != nil && file.Exists(f.Value.String()) {
		return f.Value.String(), nil
	}

	// ENV
	if p := os.Getenv("CONFIG_FILE"); p != "" && file.Exists(p) {
		return p, nil
	}

	return "", errors.New("config file not found")
}
