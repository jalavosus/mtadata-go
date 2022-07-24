package utils

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

var (
	ErrEnumNotString = errors.New("gql enums must be strings")
)

type InvalidEnumError struct {
	val any
}

func invalidEnumError(val any) *InvalidEnumError {
	return &InvalidEnumError{val: val}
}

func (e *InvalidEnumError) Error() string {
	return fmt.Sprintf("value '%[1]v' invalid for enum type %[1]T", e.val)
}

type GqlEnum interface {
	~int32
	IsValid() bool
}

func SerializeGQL(s string, w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(s))
}

func DeserializeGQL[T GqlEnum](val any, out *T, fromString IotaFromStringFn[T]) error {
	str, ok := val.(string)
	if !ok {
		return ErrEnumNotString
	}

	enumVal := fromString(str)
	if !enumVal.IsValid() {
		return invalidEnumError(enumVal)
	}

	*out = enumVal

	return nil
}

// func ParseDbFieldsFromGqlQuery(ctx context.Context) []string {
//
// }
