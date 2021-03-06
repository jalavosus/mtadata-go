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

func DeserializeEnum[T ~int32](data []byte, out *T, serializeType SerializeType, fn IotaFromStringFn[T]) error {
	res, err := DeserializeStringEnum(data, serializeType)
	if err != nil {
		return err
	}

	*out = fn(res)

	return nil
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

func DeserializeStringEnum(b []byte, serializeType SerializeType) (string, error) {
	var (
		buf bytes.Buffer
		err error
		res string
	)

	if _, err = buf.Write(b); err != nil {
		return res, err
	}

	switch serializeType {
	case SerializeJson:
		err = json.NewDecoder(&buf).Decode(&res)
	case SerializeYaml:
		err = yaml.NewDecoder(&buf).Decode(&res)
	}

	return res, nil
}
