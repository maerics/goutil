package goutil

import (
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

func Boolp(b bool) *bool          { return &b }
func Stringp(s string) *string    { return &s }
func Intp(x int) *int             { return &x }
func Int8p(x int8) *int8          { return &x }
func Int16p(x int16) *int16       { return &x }
func Int32p(x int32) *int32       { return &x }
func Int64p(x int64) *int64       { return &x }
func Uintp(x uint) *uint          { return &x }
func Uint8p(x uint8) *uint8       { return &x }
func Uint16p(x uint16) *uint16    { return &x }
func Uint32p(x uint32) *uint32    { return &x }
func Uint64p(x uint64) *uint64    { return &x }
func Bytep(b byte) *byte          { return &b }
func Runep(r rune) *rune          { return &r }
func Float32p(f float32) *float32 { return &f }
func Float64p(f float64) *float64 { return &f }

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

func Keys[T comparable, U any](m map[T]U) []T {
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
