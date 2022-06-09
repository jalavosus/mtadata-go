package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

func EnumFromString[T fmt.Stringer](s string, validValues []T, unknown T) T {
	var (
		iotaVal T
		ok      bool
	)

	for _, val := range validValues {
		if strings.ToUpper(s) == StringerToUpper(val) {
			iotaVal = val
			ok = true
			break
		}
	}

	if !ok {
		iotaVal = unknown
	}

	return iotaVal
}

func EnumToDbValue(val fmt.Stringer) driver.Value {
	s := val.String()
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, " ", "_")

	return s
}

func DbValueToEnum[T fmt.Stringer](value string, validValues []T, unknown T) T {
	value = strings.ReplaceAll(value, "_", " ")
	return EnumFromString[T](value, validValues, unknown)
}
