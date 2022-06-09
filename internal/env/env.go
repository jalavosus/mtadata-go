package env

import (
	"os"

	"github.com/jalavosus/mtadata/internal/utils"
)

type parseFn[T any] func(s string) T

func fromEnv(envKey string) (string, bool) {
	return os.LookupEnv(envKey)
}

func fromEnvFallback[T any](envKey string, fallback T, parse parseFn[T]) T {
	if val, ok := fromEnv(envKey); ok {
		return parse(val)
	}

	return fallback
}

func StringFromEnv(envKey string, fallback string) string {
	if val, ok := fromEnv(envKey); ok {
		return val
	}

	return fallback
}

func Int64FromEnv(envKey string, fallback int64) int64 {
	return fromEnvFallback(envKey, fallback, utils.ParseInt64)
}

func IntFromEnv(envKey string, fallback int) int {
	return int(Int64FromEnv(envKey, int64(fallback)))
}

func Uint64FromEnv(envKey string, fallback uint64) uint64 {
	return fromEnvFallback(envKey, fallback, utils.ParseUint64)
}

func BoolFromEnv(envKey string, fallback bool) bool {
	return fromEnvFallback(envKey, fallback, utils.ParseBool)
}

func Float64FromEnv(envKey string, fallback float64) float64 {
	return fromEnvFallback(envKey, fallback, utils.ParseFloat64)
}
