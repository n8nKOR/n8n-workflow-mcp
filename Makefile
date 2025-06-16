# Makefile
.PHONY: build build-linux build-windows test clean run help fmt lint mod-tidy build-standalone build-standalone-linux build-standalone-windows

# Default target
.DEFAULT_GOAL := help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Build info
BINARY_NAME=n8n-mcp-server
BINARY_STANDALONE=$(BINARY_NAME)
BINARY_STANDALONE_UNIX=$(BINARY_NAME)_unix
BINARY_STANDALONE_WIN=$(BINARY_NAME).exe
BUILD_DIR=build

## build: Build the standalone application (default)
build: build-standalone

## build-linux: Build the standalone application for linux
build-linux: build-standalone-linux

## build-windows: Build the standalone application for Windows
build-windows: build-standalone-windows

## build-standalone: Build the standalone application with embedded data (local OS)
build-standalone:
	@echo "Building $(BINARY_STANDALONE) with embedded data..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_STANDALONE) ./cmd/n8n-mcp-server

## build-standalone-linux: Build the standalone application for Linux with embedded data
build-standalone-linux:
	@echo "Building $(BINARY_STANDALONE_UNIX) for Linux with embedded data..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_STANDALONE_UNIX) ./cmd/n8n-mcp-server

## build-standalone-windows: Build the standalone application for Windows with embedded data
build-standalone-windows:
	@echo "Building $(BINARY_STANDALONE_WIN) for Windows with embedded data..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_STANDALONE_WIN) ./cmd/n8n-mcp-server

## test: Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

## clean: Clean build files
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

## run: Run the standalone application
run: build-standalone
	@echo "Running $(BINARY_STANDALONE) in standalone mode..."
	$(BUILD_DIR)/$(BINARY_STANDALONE) -standalone

## run-standalone: Run the standalone application (alias for run)
run-standalone: run

## fmt: Format Go code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -s -w .

## lint: Run linter
lint:
	@echo "Running linter..."
	$(GOLINT) run

## mod-tidy: Tidy go modules
mod-tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

## mod-download: Download go modules
mod-download:
	@echo "Downloading modules..."
	$(GOMOD) download

## help: Show this help message
help:
	@echo "Available commands:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
