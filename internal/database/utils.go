package database

import (
	"github.com/jalavosus/mtadata/internal/utils"
)

func checkValid[T comparable](val *T, invalid T) (res T, ok bool) {
	if v, ptrOk := utils.FromPointer[T](val); ptrOk && v != invalid {
		res = v
		ok = true
	}

	return
}

func checkValidParam[T any](val QueryParam[T]) (res any, ok bool) {
	if v, ptrOk := utils.FromPointer(val.Arg()); ptrOk && !val.Invalid() {
		res = v
		ok = true
	}

	return
}
