BINARY := lfx

.PHONY: build test lint

build:
	go build -o bin/$(BINARY) ./cmd/lfx

test:
	go test ./...

lint:
	gofmt -w cmd internal
