.PHONY: build run test clean generate-certs

# Build variables
BINARY_NAME=gochat
GO=go

build:
	$(GO) build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME) -addr=localhost:8080 -cert=server.crt -key=server.key

test:
	$(GO) test -v ./...

clean:
	rm -f $(BINARY_NAME)
	rm -f server.crt server.key

generate-certs:
	openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"

# Install dependencies
deps:
	$(GO) mod tidy

# Run with race detection
race:
	$(GO) run -race main.go

# Build for different platforms
build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-linux-amd64 main.go

build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-windows-amd64.exe main.go

build-mac:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-darwin-amd64 main.go

# Build all platforms
build-all: build-linux build-windows build-mac

# Default target
all: clean generate-certs build test
