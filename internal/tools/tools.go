//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/gogo/protobuf/protoc-gen-gofast"
	_ "github.com/golang/protobuf"
	_ "google.golang.org/genproto"
)
