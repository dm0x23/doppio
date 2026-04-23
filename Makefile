BINARY_NAME=dop
MAIN_PATH=./cmd/dop
GO=go

.PHONY: build install uninstall clean test fmt

build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)

install:
	$(GO) install $(MAIN_PATH)

uninstall:
	rm -f $(shell go env GOPATH)/bin/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...
