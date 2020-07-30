export VERSION := v0.3.0
export BINARY_NAME := gojson
export PACKAGE_NAME := gojson

# Go releated variables
# Sources

BASE_DIR := $(shell pwd)
GOBIN := $(BASE_DIR)/out
MAINGO := $(BASE_DIR)/cmd/gojson/gojson.go

PLATFORMS := $(shell [[ -z "${PLATFORM}" ]] && echo windows linux darwin || echo ${PLATFORM} )
# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-s -w -X gojson/internal/common.Version=$(VERSION)"

# Make it silent
MAKEFLAGS += --silent

build: clean $(PLATFORMS)
	@echo "> Buid finished"

%:
	@echo "> Buiding gojson for $@"
	@mkdir -p $(GOBIN)/$@
	CGO_ENABLED=0 GOOS=$@ GOARCH=amd64 go build -trimpath $(LDFLAGS) -o $(GOBIN)/$@ $(MAINGO)
	cd out/$@ && zip gojson-$@-amd64-$(VERSION).zip *

clean:
	@echo "> Cleaning builded files..."
	@-rm -rf $(GOBIN)/* 2> /dev/null

cleanf: clean
	@echo "> Cleaning cache..."
	go clean

test:
	@echo "> Test started"
	@CGO_ENABLED=0 go test -v ./...

test-cover:
	@echo "> Coverage started"
	@CGO_ENABLED=0 go test -v ./... -coverprofile=./out/cover.out

coveralls: test-cover
	@goveralls -coverprofile=./out/cover.out -service=drone.io

.PHONY: clean cleanf test test-cover coveralls
