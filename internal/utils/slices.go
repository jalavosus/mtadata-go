package utils

type MapSliceFn[T, U any] func(T) U

func MapSlice[T, U any](data []T, fn MapSliceFn[T, U]) (mappedSlice []U) {
	mappedSlice = make([]U, len(data))

	for i, d := range data {
		mappedSlice[i] = fn(d)
	}

	return
}

func SliceContains[T comparable](data []T, val T) bool {
	for i := range data {
		if data[i] == val {
			return true
		}
	}

	return false
}
