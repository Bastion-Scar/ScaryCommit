BINARY_LINUX=sco-linux
BINARY_WINDOWS=sco-windows.exe
BINARY_MACOS=sco-macos

.PHONY: all build test lint clean docker

all: build

build:
	@echo "Building all binaries..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_LINUX) .
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_WINDOWS) .
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_MACOS) .

test:
	@echo "Running tests..."
	go test ./...

lint:
	@echo "Linting..."
	golangci-lint run

docker:
	@echo "Building Docker image..."
	docker build -t scarycommit-builder .

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_LINUX) $(BINARY_WINDOWS) $(BINARY_MACOS)
