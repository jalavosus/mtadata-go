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
	return enumToDbValue(val)
}

func enumToDbValue(val fmt.Stringer) string {
	s := val.String()
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, " ", "_")

	return s
}

func DbValueToEnum[T fmt.Stringer](value string, validValues []T, unknown T) T {
	value = strings.ReplaceAll(value, "_", " ")
	return EnumFromString[T](value, validValues, unknown)
}

func MakeCreateEnumTypeCommand[T fmt.Stringer](validValues []T, typeName string) string {
	var enumValues = make([]string, len(validValues))
	for i, val := range validValues {
		enumValues[i] = fmt.Sprintf(`'%[1]s'`, enumToDbValue(val))
	}

	joinedValues := strings.Join(enumValues, ",")

	return fmt.Sprintf("CREATE TYPE public.%[1]s AS ENUM (%[2]s);", typeName, joinedValues)
}
