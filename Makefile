SRC = $(find . -name '*.go')

.PHONY: all
all: test fmt lint

.PHONY: test
test: ${SRC}
	go test ./...

fmt: ${SRC}
	go fmt ./...

lint: ${SRC}
	golint ./...
