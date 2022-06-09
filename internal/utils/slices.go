package utils

type MapSliceFn[T, U any] func(T) U

func MapSlice[T, U any](data []T, fn MapSliceFn[T, U]) (mappedSlice []U) {
	mappedSlice = make([]U, len(data))

	for i, d := range data {
		mappedSlice[i] = fn(d)
	}

	return
}
