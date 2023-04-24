#
SHELL := /bin/bash

# GENERATING
.PHONY: generate
generate: make generate

.PHONY: generate
generate: pkg/ergast/server.gen.go

pkg/ergast/server.gen.go: pkg/ergast/api.yaml
pkg/ergast/server.gen.go: pkg/ergast/codegen.yaml
pkg/ergast/server.gen.go: pkg/ergast/server.go
	go generate pkg/ergast/server.go


# CLEANING
.PHONY: clean
clean: clean-gogen 

.PHONY: clean-gogen
clean-gogen:
	rm -f pkg/**/*.gen.go
