.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test

test:
	go test -v -race --timeout 30s ./... 	

.PHONY: lint
lint: 
	go vet ./...
	
.DEFAULT_GOAL := build		