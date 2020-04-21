export VERSION := v0.2.1
export BINARY_NAME := gojson
export PACKAGE_NAME := gojson

# Go releated variables
# Sources

BASE_DIR := $(shell pwd)
GOBIN := $(BASE_DIR)/out
MAINGO := $(BASE_DIR)/cmd/gojson/gojson.go

PLATFORMS := windows linux darwin
# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X gojson/internal/common.Version=$(VERSION)"

# Make it silent
MAKEFLAGS += --silent

build: clean $(PLATFORMS)
	@echo "> Buid finished"

%:
	@echo "> Buiding gojson for $@"
	@mkdir -p $(GOBIN)/$@
	GOOS=$@ GOARCH=amd64 go build $(LDFLAGS) -o $(GOBIN)/$@ $(MAINGO)
	cd out/$@ && zip gojson-$@-amd64-$(VERSION).zip *

clean:
	@echo "> Cleaning builded files..."
	@-rm -rf $(GOBIN)/* 2> /dev/null
