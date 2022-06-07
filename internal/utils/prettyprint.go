package utils

import (
	"bytes"

	"github.com/k0kubun/pp"
)

func init() {
	colorScheme := pp.ColorScheme{
		FieldName: pp.Green,
		String:    pp.Blue,
	}

	pp.SetColorScheme(colorScheme)
	pp.PrintMapTypes = false
}

func PrettyPrintStruct(val any, pkgPrefix string, structNames ...string) string {
	out := new(bytes.Buffer)
	_, _ = pp.Fprint(out, val)

	var b = out.Bytes()

	for _, rep := range structNames {
		repStr := []byte(pkgPrefix + ".\x1b[32m" + rep + "\x1b[0m")
		b = bytes.ReplaceAll(b, repStr, []byte(""))
	}

	return string(b)
}
