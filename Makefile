GOFMT=gofmt -w
DEPS=$(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES := $(shell go list ./...)

# Prettify output
UNAME_S	 := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	RESET = $(shell echo -e "\033[0m")
	GREEN = $(shell echo -e "\033[32;01m")
	ERROR = $(shell echo -e "\033[31;01m")
	WARN  = $(shell echo -e "\033[33;01m")
endif
ifeq ($(UNAME_S),Darwin)
	RESET := $(shell echo "\033[0m")
	GREEN := $(shell echo "\033[32;01m")
	ERROR := $(shell echo "\033[31;01m")
	WARN  := $(shell echo "\033[33;01m")
endif

default: build

dep:	## Install dependencies needed for this application.
	@echo "$(GREEN)>>> Installing dependencies$(RESET)"
	@go get -u -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d

proto:	## Generate the protocol buffers with protoc.
	@echo "$(GREEN)>>> Generating protocol buffers$(RESET)"
	@if ! which protoc > /dev/null; then \
		echo "$(WARN)Error: protoc not found$(GREEN)" >&2; \
		exit 1; \
	fi
	go get -u -v github.com/golang/protobuf/protoc-gen-go
	@for dir in $$(git ls-files '*.proto' | xargs -n1 dirname | uniq); do \
		protoc -I $$dir --go_out=plugins=grpc:$$dir $$dir/*.proto; \
	done

format:	## Format source files with gofmt.
	@echo "$(GREEN)>>> Formatting$(RESET)"
	$(foreach ENTRY,$(PACKAGES),$(GOFMT) $(GOPATH)/src/$(ENTRY);)

build:	## Build binaries.
	@echo "$(GREEN)>>> Building$(RESET)"
	go build -o ./ramify        ./cmd/client
	go build -o ./ramify-post   ./cmd/post
	go build -o ./ramify-api    ./cmd/api
	go build -o ./ramify-auth   ./cmd/auth

clean:  ## Remove previous builds binaries.
	@echo "$(GREEN)>>> Cleaning$(RESET)"
	go clean -i -r -x
	rm ./ramify && rm ./ramify-post && rm ./ramify-auth && rm ./ramify-api

install:	## Install binaries to $GOPATH/bin.
	@echo "$(GREEN)>>> Installing$(RESET)"
	install ./ramify-post	$(GOPATH)/bin
	install ./ramify 	$(GOPATH)/bin
	install ./ramify-api	$(GOPATH)/bin
	install ./ramify-auth	$(GOPATH)/bin

lint:	## Run linter on source directories.
	@echo "$(GREEN)>>> Linting$(RESET)"
	$(GOPATH)/bin/golint ./cmd/client
	$(GOPATH)/bin/golint ./cmd/post	
	$(GOPATH)/bin/golint ./cmd/auth
	$(GOPATH)/bin/golint ./cmd/api

vet:	## Run go vet on source directories.
	go vet ./cmd/ramify/
	go vet ./cmd/post/	
	go vet ./cmd/api/
	go vet ./cmd/auth/

all: format dep proto vet build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
