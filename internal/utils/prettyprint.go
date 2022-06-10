package utils

import (
	"bytes"

	"github.com/k0kubun/pp/v3"
)

func init() {
	colorScheme := pp.ColorScheme{
		FieldName: pp.Green,
		String:    pp.Blue,
	}

	pp.SetColorScheme(colorScheme)
	pp.PrintMapTypes = false
}

type PrettyPrintParam struct {
	PkgPrefix string
	TypeNames []string
}

func NewPrettyPrintParam(pkgPrefix string, typeNames ...string) PrettyPrintParam {
	return PrettyPrintParam{
		PkgPrefix: pkgPrefix,
		TypeNames: typeNames,
	}
}

func PrettyPrintStruct(val any, params ...PrettyPrintParam) string {
	out := new(bytes.Buffer)
	_, _ = pp.Fprint(out, val)

	var b = out.Bytes()

	for _, p := range params {
		for _, typeName := range p.TypeNames {
			repStr := []byte(p.PkgPrefix + ".\x1b[32m" + typeName + "\x1b[0m")
			b = bytes.ReplaceAll(b, repStr, []byte(""))
		}
	}

	return string(b)
}
