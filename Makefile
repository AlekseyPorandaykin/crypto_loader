HOME_PATH := $(shell pwd)

BIN := "./bin/crypto_loader"
VERSION :=$(shell date)

build:
	go build -o=$(BIN) -ldflags="-X 'main.version=${VERSION}' -X 'main.homeDir=${HOME_PATH}'" .

linters:
	go vet .
	gofmt -w .
	goimports -w .
	gci write /app
	gofumpt -l -w /app
	golangci-lint run ./...
	gofmt -s -l $(git ls-files '*.go')

.PHONY: build run build-img run-img version test lint
