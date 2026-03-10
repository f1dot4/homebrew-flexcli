VERSION ?= $(shell git describe --tags --always --dirty)
LDFLAGS := -X main.Version=$(VERSION)

.PHONY: build clean

build:
	mkdir -p bin
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o bin/flexcli-mac ./cmd/flexcli/
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o bin/flexcli-linux ./cmd/flexcli/

clean:
	rm -rf bin/
