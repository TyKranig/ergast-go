#
SHELL := /bin/bash

# GENERATING
.PHONY: generate
generate: pkg/ergast/server.gen.go

pkg/ergast/server.gen.go: pkg/ergast/api.yaml
pkg/ergast/server.gen.go: pkg/ergast/codegen.yaml
pkg/ergast/server.gen.go: pkg/ergast/server.go
	go generate pkg/ergast/server.go


# LINTING
.PHONY: format
format:
	go fmt ./...

.PHONY: lint-go
lint-go: generate
	@echo "Linting Go files"
	golangci-lint run --verbose --timeout 5m

.PHONY: lint-oas
lint-oas:
	@echo "Linting OpenAPI Specs"
	spectral lint pkg/ergast/api.yaml

.PHONY: lint
lint: lint-go
lint: lint-oas

# CLEANING
.PHONY: clean
clean: clean-gogen 

.PHONY: clean-gogen
clean-gogen:
	rm -f pkg/**/*.gen.go
