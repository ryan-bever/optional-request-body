package main

import (
	_ "embed"
)

//go:embed openapi.yaml
var OpenapiYAML string

//go:generate ./bin/oapi-codegen --config=config-server.yaml ./openapi.yaml
