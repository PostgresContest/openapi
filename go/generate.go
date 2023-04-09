package _go

import _ "github.com/ogen-go/ogen"

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target gen/v1 -package v1 --clean ../api/v1/openapi.yaml
