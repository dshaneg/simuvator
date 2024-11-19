BINDIR := $(CURDIR)/bin
BINNAME ?= simuvator

# go option
PKG := ./...
TAGS :=
TESTS := .
TESTFLAGS :=

# Rebuild the binary if any of these files change
SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

# Required for globs to work correctly
SHELL = /usr/bin/env bash

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA = $(shell git rev-parse --short HEAD)
GIT_TAG = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

# -------
# build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	go build $(GOFLAGS) -o '$(BINDIR)'/$(BINNAME) ./cmd/simuvator

# -------
# test

.PHONY: test
test:
	go test -v ./...

# -------

.PHONY: run
run:
	go run ./cmd/simuvator

.PHONY: clean
clean:
	@rm -rf '$(BINDIR)'

.PHONY: info
info:
	@echo "Git Tag:        $(GIT_TAG)"
	@echo "Git Commit:     $(GIT_COMMIT)"
	@echo "Git Sha:        $(GIT_SHA)"
	@echo "Git Tree State: $(GIT_DIRTY)"
