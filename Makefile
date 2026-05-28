BINARY_NAME=dop
MAIN_PATH=./cmd/dop
GO=go

# Version is taken from the latest git tag, fallback to "dev"
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: build install uninstall clean test fmt build-all clean-release

# Local development build (with version info, linked dynamically for speed)
build:
	$(GO) build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) $(MAIN_PATH)

install:
	$(GO) install -ldflags "$(LDFLAGS)" $(MAIN_PATH)

uninstall:
	rm -f $(shell go env GOPATH)/bin/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...

# ---- Release binaries (static, cross-compiled) ----
build-all: clean-release
	@echo "Building static release binaries (version $(VERSION))"
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o release/dop-linux-amd64      $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o release/dop-linux-arm64      $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o release/dop-darwin-amd64     $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o release/dop-darwin-arm64     $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o release/dop-windows-amd64.exe $(MAIN_PATH)
	@echo "Binaries placed in ./release/"

clean-release:
	rm -rf release
