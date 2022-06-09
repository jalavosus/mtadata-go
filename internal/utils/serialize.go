package utils

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/goccy/go-yaml"
)

type (
	SerializeType           uint8
	IotaFromStringFn[T any] func(s string) T
)

const (
	SerializeJson SerializeType = iota
	SerializeYaml
	SerializeBytes
)

func SerializeEnum(data fmt.Stringer, serializeType SerializeType) ([]byte, error) {
	return SerializeString(data.String(), serializeType)
}

func DeserializeEnum[T any](data []byte, fromStringFn IotaFromStringFn[T]) T {
	s := fmt.Sprintf("%x", data)
	return fromStringFn(s)
}

func SerializeString(s string, serializeType SerializeType) ([]byte, error) {
	var (
		buf bytes.Buffer
		err error
	)

	switch serializeType {
	case SerializeJson:
		err = json.NewEncoder(&buf).Encode(s)
	case SerializeYaml:
		err = yaml.NewEncoder(&buf).Encode(s)
	case SerializeBytes:
		_, err = buf.WriteString(s)
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
