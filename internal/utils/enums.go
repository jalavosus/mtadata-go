package utils

import (
	"database/sql/driver"
	"strings"

	"github.com/jalavosus/mtadata/models/enums"
)

func EnumFromString[T enums.StringEnum](s string, validValues []T, unknown T) T {
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

func EnumToDbValue(val enums.StringEnum) driver.Value {
	s := val.String()
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, " ", "_")

	return s
}

func DbValueToEnum[T enums.StringEnum](value string, validValues []T, unknown T) T {
	value = strings.ReplaceAll(value, "_", " ")
	return EnumFromString[T](value, validValues, unknown)
}
