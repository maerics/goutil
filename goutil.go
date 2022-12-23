package goutil

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	log "github.com/maerics/golog"
	"golang.org/x/exp/constraints"
)

type A []any
type M map[string]any

func MustEnv(varname string) string {
	value := strings.TrimSpace(os.Getenv(varname))
	if value == "" {
		log.Fatalf("the environment variable %q must be set", varname)
	}
	return value
}

func MustJson(x any, pretty ...bool) string {
	if len(pretty) > 0 && pretty[0] {
		return string(log.Must1(json.MarshalIndent(x, "", "  ")))
	}
	return string(log.Must1(json.Marshal(x)))
}

func Keys[T comparable](m map[T]any) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Sort[T constraints.Ordered](xs []T) []T {
	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	return xs
}
