package utils

import (
	"strconv"
)

// ParseFloat32 wraps strconv.ParseFloat,
// panicking if strconv.ParseFloat returns an error.
func ParseFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		panic(err)
	}

	return float32(f)
}

// ParseFloat64 wraps strconv.ParseFloat,
// panicking if strconv.ParseFloat returns an error.
func ParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	return f
}

// ParseInt wraps strconv.ParseInt,
// panicking if strconv.ParseInt returns an error.
func ParseInt(s string) int {
	return int(ParseInt64(s))
}

// ParseInt64 wraps strconv.ParseInt,
// panicking if strconv.ParseInt returns an error.
func ParseInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return n
}

// ParseUint64 wraps strconv.ParseUint,
// panicking if strconv.ParseUint returns an error.
func ParseUint64(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return n
}

// ParseBool wraps strconv.ParseBool,
// panicking if strconv.ParseBool returns an error.
func ParseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}

	return b
}
