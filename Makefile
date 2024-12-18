# Setting SHELL to bash allows bash commands to be executed by recipes.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL := /usr/bin/env bash -o pipefail
.SHELLFLAGS := -ec

include .bingo/Variables.mk

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run $(GOLANGCI_LINT_ARGS)

.PHONY: unit
unit:
	go test ./... -cover -coverprofile unit.out

.PHONY: build
build:
	go build -o bin/crd-diff main.go
