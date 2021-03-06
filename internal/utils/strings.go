package utils

import (
	"fmt"
	"strings"
)

func String(s fmt.Stringer) string {
	return s.String()
}

func ToStringSlice[T fmt.Stringer](data []T) (res []string) {
	res = make([]string, len(data))

	for i, s := range data {
		res[i] = s.String()
	}

	return
}

func StringerToUpper(s fmt.Stringer) string {
	return strings.ToUpper(s.String())
}

func TrimParens(s string) string {
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")

	return s
}

func TrimWhitespaceSlice(data []string) (res []string) {
	res = make([]string, len(data))

	for i, d := range data {
		res[i] = strings.TrimSpace(d)
	}

	return
}
