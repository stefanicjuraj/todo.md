.PHONY: build install clean test help

BINARY_NAME=todo
VERSION?=0.1.0
BUILD_DIR=bin
GO_FILES=$(shell find . -type f -name '*.go' -not -path './vendor/*')

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

install: build ## Install the binary to /usr/local/bin (requires sudo)
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "Installed successfully!"

install-user: build ## Install the binary to ~/go/bin (no sudo required)
	@echo "Installing $(BINARY_NAME) to ~/go/bin..."
	@mkdir -p ~/go/bin
	@cp $(BUILD_DIR)/$(BINARY_NAME) ~/go/bin/$(BINARY_NAME)
	@echo "Installed successfully! Make sure ~/go/bin is in your PATH."

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

test: ## Run tests
	@go test -v ./...

fmt: ## Format the code
	@go fmt ./...

vet: ## Run go vet
	@go vet ./...

lint: fmt vet ## Run linters

deps: ## Download dependencies
	@go mod download
	@go mod tidy

.DEFAULT_GOAL := help

