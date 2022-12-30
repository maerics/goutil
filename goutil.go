package goutil

import (
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

type A []any
type M map[string]any

func Getenv(name, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func MustEnv(varname string) string {
	value := strings.TrimSpace(os.Getenv(varname))
	if value == "" {
		log.Fatalf("the environment variable %q must be set", varname)
	}
	return value
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
