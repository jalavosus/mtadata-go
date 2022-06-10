package utils

func ToPointer[T any](val T) *T {
	return &val
}

func FromPointer[T any](val *T) (res T, ok bool) {
	if val != nil {
		res = *val
		ok = true
	}

	return
}
