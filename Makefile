PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

.PHONY: all build install clean test lint docker

all: build install
build:
	@go build -v ./...
install:
	@go install
clean:
	@rm $(GOPATH)/bin/goproc
test:
	@go test -v ./...
lint:
	@golangci-lint run --enable-all
docker:
	@docker build -t goproc .
