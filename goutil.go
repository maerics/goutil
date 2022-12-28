package goutil

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"

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

func Getenv(name, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func MustJson(x any, pretty ...bool) string {
	var bs []byte
	var err error
	if len(pretty) > 0 && pretty[0] {
		bs, err = json.MarshalIndent(x, "", "  ")
	} else {
		bs, err = json.Marshal(x)
	}
	if err != nil {
		panic(err)
	}
	return string(bs)
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
