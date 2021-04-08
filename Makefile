-include .env
.DEFAULT_GOAL:=help
SHELL:=/bin/sh

.PHONY: build clean run setup help

VERSION := $(or $(shell git describe --tags 2>/dev/null), Unknown)
BUILD := $(or $(shell git rev-parse --short HEAD), Unknown)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
MODULE = $(shell env GO111MODULE=on $(GO) list -m)
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(filter-out tools.go, $(wildcard *.go))

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags '-w -X=main.Version=$(VERSION) -X=main.Build=$(BUILD)'

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

export GO111MODULE=on
undefine GOPATH

help: ##Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

clean: .go-clean ##Cleanup the project files

build: .go-get .go-build ##Build the project

.go-build:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

.go-generate:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)

.go-get:
	GOBIN=$(GOBIN) go get $(get)
	GOBIN=$(GOBIN) go mod vendor

.go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

.go-clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean